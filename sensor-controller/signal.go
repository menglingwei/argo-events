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

package sensor_controller

import (
	"github.com/argoproj/argo-events/pkg/apis/sensor/v1alpha1"
	"github.com/argoproj/argo-events/pkg/event"
)

func (sc *sensorCtx) processSignal(name string, event event.Event) (*v1alpha1.NodeStatus, error) {
	sc.log.Info().Str("signal-name", name).Msg("processing signal")
	node := getNodeByName(sc.sensor, name)
	if node == nil {
		node = sc.initializeNode(signal, v1alpha1.NodeTypeSignal, v1alpha1.NodePhaseNew)
		// we can skip active state.
	}
	node.LatestEvent = &v1alpha1.EventWrapper{
		Event: event,
		Seen: true,
	}
	return soc.markNodePhase(signal, v1alpha1.NodePhaseComplete, "signal processing completed"), nil
}
