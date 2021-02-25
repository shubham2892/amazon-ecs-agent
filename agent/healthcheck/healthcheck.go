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

package healthcheck

//go:generate mockgen -destination=mock/$GOFILE -copyright_file=../../scripts/copyright_file github.com/aws/amazon-ecs-agent/agent/healthcheck Healthcheck

import (
	"context"
	"time"

	"github.com/cihub/seelog"
)

// heartbeats with relative frequencies in ratio 1:2:4
// write: 15sec, read: 30sec, healthcheck: 1min
const (
	HeartbeatWriteInterval = 4* time.Second
	HeartbeatReadInterval  = 2 * HeartbeatWriteInterval
	HealthcheckInterval    = 4 * HeartbeatWriteInterval
)

// var should contain any error types specific to Healthcheck
//var ()

// Healthcheck defines methods to be implemented by the healthcheck struct.
//type Healtcheck interface {
//	NewNamedHeartbeat(hbName string) (*NamedHeartbeat)
//	StartHealthcheckProcess(ctx context.Context)
//	GetAgentOverallHealthStatus() (bool)
//}

type NamedHeartbeat struct {
	HeartbeatName    string
	HeartbeatChannel chan interface{}
}

// overallHealthStatus will be a specific health type of CONNECTED/IMPAIRED/DISCONNECTED ??
// or will it be a bool?  Disconnected would not be surfaced by agent
type Healthcheck struct {
	ctx                         context.Context
	heartbeatScorecard          map[string]bool
	namedHeartbeats             []*NamedHeartbeat
	aggregateHeartbeatIsHealthy bool
	containerRuntimeIsHealthy   bool
}

// NewHealthcheck creates a new instance of the Healthcheck object.
func NewHealthcheck(ctx context.Context) *Healthcheck {
	derivedCtx, _ := context.WithCancel(ctx)
	return &Healthcheck{
		ctx:                         derivedCtx,
		heartbeatScorecard:           make(map[string]bool),
		namedHeartbeats:             make([]*NamedHeartbeat, 0),
		aggregateHeartbeatIsHealthy: false,
		containerRuntimeIsHealthy:   false,
	}
}

func (hc *Healthcheck) NewNamedHeartbeat(hbName string) *NamedHeartbeat {
	namedHeartbeat := &NamedHeartbeat{
		HeartbeatName:    hbName,
		HeartbeatChannel: make(chan interface{}),
	}
	hc.namedHeartbeats = append(hc.namedHeartbeats, namedHeartbeat)
	return namedHeartbeat
}

func (hc *Healthcheck) StartHealthcheckProcess(ctx context.Context) {
	heartbeatReadPulse := time.Tick(HeartbeatReadInterval)
	healthcheckPulse := time.Tick(HealthcheckInterval)
	for {
		// non-blocking reads of heartbeats.
		select {
		case <-heartbeatReadPulse:
			for _, hb := range hc.namedHeartbeats {
				seelog.Infof("checking namedHeartbeat: %s", hb.HeartbeatName)
				select {
				case _, ok := <-hb.HeartbeatChannel:
					seelog.Infof("in namedHeartbeat channel, ok is: %v", ok)
					//if ok == false {
					//		continue
					//}
					seelog.Infof("namedHeartbeat: %s is healthy", hb.HeartbeatName)
					hc.heartbeatScorecard[hb.HeartbeatName] = true
					seelog.Infof("heartbeatScorecard: %v", hc.heartbeatScorecard)
				//default:
				}
			}
		case <-healthcheckPulse:
			seelog.Infof("healthcheck pulse received")
			// should we default to healthy false?
			healthy := true
			for _, value := range hc.heartbeatScorecard {
				if value == false {
					healthy = false
					break
				}
			}
			// set the healthstatus
			hc.aggregateHeartbeatIsHealthy = healthy
			seelog.Infof("healthcheck status: %v", healthy)
			// reset hc.heartbeatScorecard and health flag each Healthcheck interval
			for key, _ := range hc.heartbeatScorecard {
				hc.heartbeatScorecard[key] = false
			}
			healthy = true
		}
	}
}

func (hc *Healthcheck) GetAgentOverallHealthStatus() bool {
	return hc.aggregateHeartbeatIsHealthy && hc.containerRuntimeIsHealthy
}
