// Copyright © 2020
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"knative.dev/client/pkg/kn/plugin"

	"github.com/rhuss/kn-service-log/pkg"
)

func init() {
	plugin.InternalPlugins = append(plugin.InternalPlugins, &logPlugin{})
}

type logPlugin struct {}

func (l *logPlugin) Name() string {
	return "kn-service-log"
}

func (l *logPlugin) Execute(args []string) error {
	cmd := pkg.NewLogCommand()
	return cmd.Execute()
}

func (l *logPlugin) Description() (string, error) {
	return "Print logs of pods belonging to a service", nil
}

func (l *logPlugin) CommandParts() []string {
	return []string{ "service", "log" }
}

// Path is empty because its an internal plugins
func (l *logPlugin) Path() string {
	return ""
}
