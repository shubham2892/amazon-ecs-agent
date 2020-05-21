// +build !windows
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

package stats

import (
	"fmt"

	"github.com/cihub/seelog"
	"github.com/docker/docker/api/types"
	"github.com/vishvananda/netlink"
	"github.com/containernetworking/cni/pkg/ns"
	//"github.com/aws/amazon-ecs-cni-plugins/pkg/cninswrapper"
	"github.com/pkg/errors"



)

// dockerStatsToContainerStats returns a new object of the ContainerStats object from docker stats.
func dockerStatsToContainerStats(dockerStats *types.StatsJSON) (*ContainerStats, error) {
	// The length of PercpuUsage represents the number of cores in an instance.
	if len(dockerStats.CPUStats.CPUUsage.PercpuUsage) == 0 || numCores == uint64(0) {
		seelog.Debug("Invalid container statistics reported, no cpu core usage reported")
		return nil, fmt.Errorf("Invalid container statistics reported, no cpu core usage reported")
	}

	cpuUsage := dockerStats.CPUStats.CPUUsage.TotalUsage / numCores
	memoryUsage := dockerStats.MemoryStats.Usage - dockerStats.MemoryStats.Stats["cache"]
	storageReadBytes, storageWriteBytes := getStorageStats(dockerStats)
	getDockerStats()
	networkStats := getNetworkStats(dockerStats)
	return &ContainerStats{
		cpuUsage:          cpuUsage,
		memoryUsage:       memoryUsage,
		storageReadBytes:  storageReadBytes,
		storageWriteBytes: storageWriteBytes,
		networkStats:      networkStats,
		timestamp:         dockerStats.Read,
	}, nil
}

func getStorageStats(dockerStats *types.StatsJSON) (uint64, uint64) {
	// initialize block io and loop over stats to aggregate
	if dockerStats.BlkioStats.IoServiceBytesRecursive == nil {
		seelog.Debug("Storage stats not reported for container")
		return uint64(0), uint64(0)
	}
	storageReadBytes := uint64(0)
	storageWriteBytes := uint64(0)
	for _, blockStat := range dockerStats.BlkioStats.IoServiceBytesRecursive {
		switch op := blockStat.Op; op {
		case "Read":
			storageReadBytes += blockStat.Value
		case "Write":
			storageWriteBytes += blockStat.Value
		default:
			//ignoring "Async", "Total", "Sum", etc
			continue
		}
	}
	return storageReadBytes, storageWriteBytes
}

func getDockerStats() {
	fmt.Print("Getting docker stats")
	var linksInTaskNetNS []netlink.Link
	//ns := cninswrapper.NewNS()

	err := ns.WithNetNSPath("net/ns/path", func() error {
		var linkErr error
		linksInTaskNetNS, linkErr = netlink.LinkList()
		if linkErr != nil {
			return errors.Wrap(linkErr, "failed to get network links")
		}
		return nil
	})
	if err!=nil {
		fmt.Print("ERrorer: %v", err)
	}
	//defer m.Done(err)()

	//var deviceNames []string
	for _, link := range linksInTaskNetNS {
		fmt.Print(link)
	}
}