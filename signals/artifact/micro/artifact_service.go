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

	"github.com/argoproj/argo-events/sdk"
	"github.com/argoproj/argo-events/signals/artifact"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	svc := k8s.NewService(micro.Name("artifact"), micro.Metadata(sdk.SignalMetadata))
	svc.Init()

	// stream configuration
	// TODO: make this configurable while running through github.com/micro/go-config
	// or use google/go-cloud runtimeVars
	stream, ok := os.LookupEnv("stream-signal")
	if !ok {
		stream = "nats"
	}
	streamClient := sdk.NewMicroSignalClient(stream, svc.Client())

	sdk.RegisterSignalServiceHandler(svc.Server(), sdk.NewMicroSignalServer(artifact.New(streamClient)))

	if err := svc.Run(); err != nil {
		panic(err)
	}
}
