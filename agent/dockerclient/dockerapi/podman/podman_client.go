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
	"io"
	"time"

	apicontainer "github.com/aws/amazon-ecs-agent/agent/api/container"
	apicontainerstatus "github.com/aws/amazon-ecs-agent/agent/api/container/status"
	"github.com/aws/amazon-ecs-agent/agent/config"
	"github.com/aws/amazon-ecs-agent/agent/dockerclient"
	. "github.com/aws/amazon-ecs-agent/agent/dockerclient/dockerapi"

	"github.com/docker/docker/api/types"
	dockercontainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/pkg/errors"
)

// podmanClient implements the DockerClient interface
type podmanClient struct {
}

// NewPodmanDockerClient creates a new DockerClient implemented with Podman runtime.
// TODO: implementation.
func NewPodmanDockerClient(cfg *config.Config, ctx context.Context) (DockerClient, error) {
	return &podmanClient{}, errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) SupportedVersions() []dockerclient.DockerVersion {
	return nil
}

// TODO: implementation.
func (pm *podmanClient) KnownVersions() []dockerclient.DockerVersion {
	return nil
}

// TODO: implementation.
func (pm *podmanClient) WithVersion(dockerclient.DockerVersion) DockerClient {
	return &podmanClient{}
}

// TODO: implementation.
func (pm *podmanClient) ContainerEvents(context.Context) (<-chan DockerContainerChangeEvent, error) {
	return nil, errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) PullImage(context.Context, string, *apicontainer.RegistryAuthenticationData, time.Duration) DockerContainerMetadata {
	return DockerContainerMetadata{}
}

// TODO: implementation.
func (pm *podmanClient) CreateContainer(context.Context, *dockercontainer.Config, *dockercontainer.HostConfig, string, time.Duration) DockerContainerMetadata {
	return DockerContainerMetadata{}
}

// TODO: implementation.
func (pm *podmanClient) StartContainer(context.Context, string, time.Duration) DockerContainerMetadata {
	return DockerContainerMetadata{}
}

// TODO: implementation.
func (pm *podmanClient) StopContainer(context.Context, string, time.Duration) DockerContainerMetadata {
	return DockerContainerMetadata{}
}

// TODO: implementation.
func (pm *podmanClient) DescribeContainer(context.Context, string) (apicontainerstatus.ContainerStatus, DockerContainerMetadata) {
	return apicontainerstatus.ContainerStatusNone, DockerContainerMetadata{}
}

// TODO: implementation.
func (pm *podmanClient) RemoveContainer(context.Context, string, time.Duration) error {
	return errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) InspectContainer(context.Context, string, time.Duration) (*types.ContainerJSON, error) {
	return nil, errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) ListContainers(context.Context, bool, time.Duration) ListContainersResponse {
	return ListContainersResponse{}
}

// TODO: implementation.
func (pm *podmanClient) ListImages(context.Context, time.Duration) ListImagesResponse {
	return ListImagesResponse{}
}

// TODO: implementation.
func (pm *podmanClient) CreateVolume(context.Context, string, string, map[string]string, map[string]string, time.Duration) SDKVolumeResponse {
	return SDKVolumeResponse{}
}

// TODO: implementation.
func (pm *podmanClient) InspectVolume(context.Context, string, time.Duration) SDKVolumeResponse {
	return SDKVolumeResponse{}
}

// TODO: implementation.
func (pm *podmanClient) RemoveVolume(context.Context, string, time.Duration) error {
	return errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) ListPluginsWithFilters(context.Context, bool, []string, time.Duration) ([]string, error) {
	return nil, errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) ListPlugins(context.Context, time.Duration, filters.Args) ListPluginsResponse {
	return ListPluginsResponse{}
}

// TODO: implementation.
func (pm *podmanClient) Stats(context.Context, string, time.Duration) (<-chan *types.StatsJSON, <-chan error) {
	return nil, nil
}

// TODO: implementation.
func (pm *podmanClient) Version(context.Context, time.Duration) (string, error) {
	return "", errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) APIVersion() (dockerclient.DockerVersion, error) {
	return dockerclient.Version_1_17, errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) InspectImage(string) (*types.ImageInspect, error) {
	return nil, errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) RemoveImage(context.Context, string, time.Duration) error {
	return errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) LoadImage(context.Context, io.Reader, time.Duration) error {
	return errors.New("not implemented")
}

// TODO: implementation.
func (pm *podmanClient) Info(context.Context, time.Duration) (types.Info, error) {
	return types.Info{}, errors.New("not implemented")
}