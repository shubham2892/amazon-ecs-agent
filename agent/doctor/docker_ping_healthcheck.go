// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//      http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package doctor

import (
	"context"
	"sync"
	"time"

	"github.com/aws/amazon-ecs-agent/agent/dockerclient/dockerapi"
	"github.com/cihub/seelog"
)

type dockerRuntimeHealthCheck struct {
	mu sync.RWMutex
	// Status is the container health status
	Status HealthcheckStatus `json:"Status,omitempty"`
	// Timestamp is the timestamp when container health status changed
	TimeStamp time.Time `json:"TimeStamp,omitempty"`
	// LastStatus is the last container health status
	LastStatus HealthcheckStatus `json:"lastStatus,omitempty"`
	// LastTimeStamp is the timestamp of last container health status
	LastTimeStamp time.Time `json:"LastTimeStamp,omitempty"`
	client        dockerapi.DockerClient
}

func NewDockerRuntimeHealthCheck(client dockerapi.DockerClient) *dockerRuntimeHealthCheck {
	return &dockerRuntimeHealthCheck{
		client:    client,
		Status:    HealthcheckStatusInitializing,
		TimeStamp: time.Now(),
	}
}

func (dhc *dockerRuntimeHealthCheck) RunCheck() HealthcheckStatus {
	res := dhc.client.PingDocker(context.TODO(), time.Second*2)
	healthCheck := HealthcheckStatusInitializing
	if res.Error != nil {
		seelog.Infof("Docker Ping failed with error: %v", res.Error)
		healthCheck = HealthcheckStatusImpaired
	}
	dhc.setDockerHealthCheck(healthCheck)
	return healthCheck
}

func (dhc *dockerRuntimeHealthCheck) setDockerHealthCheck(healthStatus HealthcheckStatus) {
	dhc.mu.Lock()
	defer dhc.mu.Unlock()
	dhc.LastStatus = dhc.Status
	dhc.Status = healthStatus
	dhc.LastTimeStamp = dhc.TimeStamp
	dhc.TimeStamp = time.Now()
}

func (dhc *dockerRuntimeHealthCheck) GetHealthCheckStatus() HealthcheckStatus {
	dhc.mu.RLock()
	defer dhc.mu.RUnlock()
	return dhc.Status
}

func (dhc *dockerRuntimeHealthCheck) GetHealthCheckTime() time.Time {
	dhc.mu.RLock()
	defer dhc.mu.RUnlock()
	return dhc.TimeStamp
}

func (dhc *dockerRuntimeHealthCheck) GetLastHealthCheckStatus() HealthcheckStatus {
	dhc.mu.RLock()
	defer dhc.mu.RUnlock()
	return dhc.LastStatus
}

func (dhc *dockerRuntimeHealthCheck) GetLastHealthCheckTime() time.Time {
	dhc.mu.RLock()
	defer dhc.mu.RUnlock()
	return dhc.LastTimeStamp
}
