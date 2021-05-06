// +build unit

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

package doctor

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDoctor(t *testing.T) {
	newDoctor, _ := NewDoctor()
	assert.True(t, len(newDoctor.healthchecks) == 0)
}

type testHealthcheck struct {
	testName string
}

func (thc *testHealthcheck) RunCheck() bool {
	return true
}

func (thc *testHealthcheck) GetHealthcheckStatus() HealthcheckStatus {
	return HealthcheckStatusOk
}

func (thc *testHealthcheck) GetHealthcheckTime() time.Time {
	return time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
}

func TestAddHealthcheck(t *testing.T) {
	newDoctor, _ := NewDoctor()
	assert.True(t, len(newDoctor.healthchecks) == 0)
	newTestHealthcheck := &testHealthcheck{testName: "testAddHealthcheck"}
	newDoctor.AddHealthcheck(newTestHealthcheck)
	assert.True(t, len(newDoctor.healthchecks) == 1)
}
