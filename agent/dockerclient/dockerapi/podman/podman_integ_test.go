//+build integration,podman

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

package podman

import (
	"context"
	"os"
	"testing"

	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/stretchr/testify/require"
)

func newTestClient(t *testing.T, ctx context.Context) {
	socket := os.Getenv("TEST_PODMAN_SOCKET")
	_, err := bindings.NewConnection(ctx, socket)
	require.NoError(t, err)
}

func TestNewTestClient(t *testing.T) {
	newTestClient(t, context.Background())
}