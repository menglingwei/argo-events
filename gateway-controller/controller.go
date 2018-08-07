/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gateway_controller

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	base "github.com/argoproj/argo-events"
	"github.com/argoproj/argo-events/pkg/apis/gateway/v1alpha1"
	 clientset "github.com/argoproj/argo-events/pkg/gateway-client/clientset/versioned"
)

const (
	gatewayResyncPeriod = 20 * time.Minute
)

// GatewayControllerConfig contain the configuration settings for the gateway-controller
type GatewayControllerConfig struct {
	// InstanceID is a label selector to limit the gateway-controller's watch of gateway jobs to a specific instance.
	// If omitted, the gateway-controller watches gateways that *are not* labeled with an instance id.
	InstanceID string `json:"instanceID,omitempty"`

	// Namespace is a label selector filter to limit gateway-controller's watch to specific namespace
	Namespace string `json:"namespace"`
}

// gatewayController listens for new gateways and hands off handling of each gateway on the queue to the operator
type GatewayController struct {
	// ConfigMap is the name of the config map in which to derive configuration of the contoller
	ConfigMap string
	// namespace for the config map
	ConfigMapNS string
	// Config is the gateway gateway-controller's configuration
	Config GatewayControllerConfig

	// kubernetes config and apis
	kubeConfig      *rest.Config
	kubeClientset   kubernetes.Interface
	gatewayClientset clientset.Interface

	// gateway informer and queue
	informer cache.SharedIndexInformer
	queue    workqueue.RateLimitingInterface
}

// NewGatewayController creates a new Controller
func NewGatewayController(rest *rest.Config, configMap string) *GatewayController {
	return &GatewayController{
		ConfigMap:       configMap,
		kubeConfig:      rest,
		kubeClientset:   kubernetes.NewForConfigOrDie(rest),
		clientset: clientset.NewForConfigOrDie(rest),
		queue:           workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter()),
	}
}

func (c *GatewayController) processNextItem() bool {
	// Wait until there is a new item in the queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	obj, exists, err := c.informer.GetIndexer().GetByKey(key.(string))
	if err != nil {
		log.Warnf("failed to get gateway '%s' from informer index: %+v", key, err)
		return true
	}

	if !exists {
		// this happens after gateway was deleted, but work queue still had entry in it
		return true
	}

	gateway, ok := obj.(*v1alpha1.Gateway)
	if !ok {
		log.Warnf("key '%s' in index is not a gateway", key)
		return true
	}

	ctx := newGatewayOperationCtx(gateway, c)

	err = c.handleErr(ctx.operate(), key)
	if err != nil {
		// now let's escalate the gateway
		// the context should have the most up-to-date version
		log.Infof("escalating gateway to level %s via %s message", ctx.s.Spec.Escalation.Level, ctx.s.Spec.Escalation.Message.Stream.Type)
		err := sendMessage(&ctx.s.Spec.Escalation.Message)
		if err != nil {
			log.Panicf("failed escalating gateway '%s'", key)
		}
	}

	return true
}

// handleErr checks if an error happened and make sure we will retry later
// returns an error if unable to handle the error
func (c *GatewayController) handleErr(err error, key interface{}) error {
	if err == nil {
		// Forget about the #AddRateLimited history of key on every successful sync
		// Ensure future updates for this key are not delayed because of outdated error history
		c.queue.Forget(key)
		return nil
	}

	// due to the base delay of 5ms of the DefaultControllerRateLimiter
	// requeues will happen very quickly even after a signal pod goes down
	// we want to give the signal pod a chance to come back up so we give a genorous number of retries
	if c.queue.NumRequeues(key) < 20 {
		log.Errorf("Error syncing gateway '%v': %v", key, err)

		// Re-enqueue the key rate limited. This key will be processed later again.
		c.queue.AddRateLimited(key)
		return nil
	}
	return errors.New("exceeded max requeues")
}

// Run executes the gateway-controller
func (c *GatewayController) Run(ctx context.Context, ssThreads, signalThreads int) {
	defer c.queue.ShutDown()

	log.Infof("gateway gateway-controller (version: %s) (instance: %s) starting", base.GetVersion(), c.Config.InstanceID)
	_, err := c.watchControllerConfigMap(ctx)
	if err != nil {
		log.Errorf("failed to register watch for gateway-controller config map: %v", err)
		return
	}

	c.informer = c.newgatewayInformer()
	go c.informer.Run(ctx.Done())

	if !cache.WaitForCacheSync(ctx.Done(), c.informer.HasSynced) {
		log.Panicf("timed out waiting for the caches to sync")
		return
	}

	for i := 0; i < ssThreads; i++ {
		go wait.Until(c.runWorker, time.Second, ctx.Done())
	}

	<-ctx.Done()
}

func (c *GatewayController) runWorker() {
	for c.processNextItem() {
	}
}
