// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
package handler

import (
	"context"
	"github.com/aws/amazon-ecs-agent/agent/data"
	"github.com/aws/amazon-ecs-agent/agent/engine"
	"github.com/aws/amazon-ecs-agent/agent/wsclient"
	"github.com/cihub/seelog"
	"sync"
)

// taskManifestHandler handles task manifest message for the ACS client
type dockerPingHandler struct {
	ctx                         context.Context
	taskEngine                  engine.TaskEngine
	cancel                      context.CancelFunc
	dataClient                  data.Client
	cluster                     string
	containerInstanceArn        string
	acsClient                   wsclient.ClientServer
	latestSeqNumberTaskManifest *int64
	messageId                   string
	lock                        sync.RWMutex
}

func newDockerPingHandler(ctx context.Context,
	cluster string, containerInstanceArn string, acsClient wsclient.ClientServer,
	dataClient data.Client, taskEngine engine.TaskEngine, latestSeqNumberTaskManifest *int64) dockerPingHandler {

	// Create a cancelable context from the parent context
	derivedContext, cancel := context.WithCancel(ctx)
	return dockerPingHandler{ctx: derivedContext,
		cancel:                      cancel,
		cluster:                     cluster,
		containerInstanceArn:        containerInstanceArn,
		acsClient:                   acsClient,
		taskEngine:                  taskEngine,
		dataClient:                  dataClient,
		latestSeqNumberTaskManifest: latestSeqNumberTaskManifest,
	}
}

func (dockerPingHandler *dockerPingHandler) start() {
	// Task manifest and it's ack
	//go taskManifestHandler.handleTaskManifestMessage()
	//go taskManifestHandler.sendTaskManifestMessageAck()

	// Task stop verification message and it's ack
	//go taskManifestHandler.sendTaskStopVerificationMessage()
	//go taskManifestHandler.handleTaskStopVerificationAck()

}


func (dockerPingHandler *dockerPingHandler) dockerPingMessage() {
	for {
		select {
		case <-dockerPingHandler.ctx.Done():
			return
		case message := <-taskManifestHandler.messageBufferTaskManifest:
			if err := taskManifestHandler.handleTaskManifestSingleMessage(message); err != nil {
				seelog.Warnf("Unable to handle taskManifest message [%s]: %v", message.String(), err)
			}
		}
	}
}
