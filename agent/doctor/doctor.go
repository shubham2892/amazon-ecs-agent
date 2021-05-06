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
	"github.com/cihub/seelog"
)

type Doctor struct {
	healthchecks []Healthcheck
}

func NewDoctor() (*Doctor, error) {
	return &Doctor{
		healthchecks: []Healthcheck{},
	}, nil
}

func (doc *Doctor) AddHealthcheck(healthcheck Healthcheck) {
	doc.healthchecks = append(doc.healthchecks, healthcheck)
}

func (doc *Doctor) RunHealthchecks() bool {
	//result will ultimately be an array of something...
	for _, healthcheck := range doc.healthchecks {
		res := healthcheck.RunCheck()
		seelog.Infof("healthcheck result: %v", res)
	}
	return true
}
