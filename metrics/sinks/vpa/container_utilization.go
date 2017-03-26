/*
Copyright 2017 The Kubernetes Authors.

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

package vpa

import (
	"github.com/pkg/errors"
	"k8s.io/heapster/metrics/core"
	"time"
)

type containerUtilizationSnapshot struct {
	createTime     time.Time
	scrapTime      time.Time
	containerName  string
	containerImage string
	podId          string
	memoryRequest  int64
	memoryUsage    int64
	cpuRequested   int64
	cpuUsageRate   int64
}

func newContainerUtilizationSnapshot(metricSet *core.MetricSet) (*containerUtilizationSnapshot, error) {
	err := validateMetricSet(metricSet)
	if err != nil {
		return nil, err
	}

	snapshot := containerUtilizationSnapshot{
		createTime:     metricSet.CreateTime,
		scrapTime:      metricSet.ScrapeTime,
		containerName:  metricSet.Labels[core.LabelContainerName.Key],
		containerImage: metricSet.Labels[core.LabelContainerBaseImage.Key],
		podId:          metricSet.Labels[core.LabelPodId.Key],
		memoryRequest:  metricSet.MetricValues[core.MetricMemoryRequest.Name].IntValue,
		memoryUsage:    metricSet.MetricValues[core.MetricMemoryUsage.Name].IntValue,
		cpuRequested:   metricSet.MetricValues[core.MetricCpuRequest.Name].IntValue,
		cpuUsageRate:   metricSet.MetricValues[core.MetricCpuUsageRate.Name].IntValue,
	}
	return &snapshot, nil
}

func validateMetricSet(metricSet *core.MetricSet) error {
	if metricSet.Labels == nil || metricSet.MetricValues == nil {
		return errors.New("MetricSet needs both 'Labels' and 'MetricValues' not null")
	}
	if metricSet.Labels[core.LabelContainerName.Key] == "" ||
		metricSet.Labels[core.LabelContainerBaseImage.Key] == "" ||
		metricSet.Labels[core.LabelPodId.Key] == "" {

		return errors.Errorf("MetricSet is missing one of the labels: %s, %s, %s", core.LabelContainerName.Key,
			core.LabelContainerBaseImage.Key, core.LabelPodId.Key)
	}

	return nil
}
