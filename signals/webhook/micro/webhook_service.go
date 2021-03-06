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

package main

import (
	"os"
	"strconv"

	"github.com/argoproj/argo-events/sdk"
	"github.com/argoproj/argo-events/signals/webhook"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

const (
	// EnvVarWebhookPort is the Env Var Key for the webhook port
	EnvVarWebhookPort string = "WEBHOOK_PORT"

	// DefaultWebhookPort is the default port to use if the EnvVarWebhookPort is not set
	DefaultWebhookPort int = 7070
)

func main() {
	svc := k8s.NewService(micro.Name("webhook"), micro.Metadata(sdk.SignalMetadata))
	svc.Init()

	// get the container port from container
	port := DefaultWebhookPort
	if strPort, ok := os.LookupEnv(EnvVarWebhookPort); ok {
		port, _ = strconv.Atoi(strPort)
	}

	sdk.RegisterSignalServiceHandler(svc.Server(), sdk.NewMicroSignalServer(webhook.New(port)))

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
