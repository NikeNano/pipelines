// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"time"

	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
)

const (
	ControllerAgentName string = "scheduled-workflow-controller" // ControllerAgentName is the name of the controller.
	TimeZone            string = "CRON_SCHEDULE_TIMEZONE"        // TimeZone is the name of the cron schedule env parameter
)

func getLocation() (*time.Location, error) {
	locString := common.GetStringConfigWithDefault(TimeZone, "")
	if locString == "" {
		locString = time.Now().Location().String()
	}
	return time.LoadLocation(locString)
}
