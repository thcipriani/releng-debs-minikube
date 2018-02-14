/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package config

import (
	"fmt"
	"os"
	"text/template"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/minikube/pkg/minikube/assets"
	"k8s.io/minikube/pkg/minikube/constants"
)

var addonListFormat string

type AddonListTemplate struct {
	AddonName   string
	AddonStatus string
}

var addonsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available minikube addons as well as their current statuses (enabled/disabled)",
	Long:  "Lists all available minikube addons as well as their current statuses (enabled/disabled)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Fprintln(os.Stderr, "usage: minikube addons list")
			os.Exit(1)
		}
		err := addonList()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	AddonsCmd.Flags().StringVar(&addonListFormat, "format", constants.DefaultAddonListFormat,
		`Go template format string for the addon list output.  The format for Go templates can be found here: https://golang.org/pkg/text/template/
For the list of accessible variables for the template, see the struct values here: https://godoc.org/k8s.io/minikube/cmd/minikube/cmd/config#AddonListTemplate`)
	AddonsCmd.AddCommand(addonsListCmd)
}

func stringFromStatus(addonStatus bool) string {
	if addonStatus {
		return "enabled"
	}
	return "disabled"
}

func addonList() error {
	for addonName, addonBundle := range assets.Addons {
		addonStatus, err := addonBundle.IsEnabled()
		if err != nil {
			return err
		}
		tmpl, err := template.New("list").Parse(addonListFormat)
		if err != nil {
			glog.Errorln("Error creating list template:", err)
			os.Exit(1)
		}
		listTmplt := AddonListTemplate{addonName, stringFromStatus(addonStatus)}
		err = tmpl.Execute(os.Stdout, listTmplt)
		if err != nil {
			glog.Errorln("Error executing list template:", err)
			os.Exit(1)
		}
	}
	return nil
}
