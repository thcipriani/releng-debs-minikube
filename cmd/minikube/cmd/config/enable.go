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

	"github.com/spf13/cobra"
)

var addonsEnableCmd = &cobra.Command{
	Use:   "enable ADDON_NAME",
	Short: "Enables the addon w/ADDON_NAME within minikube (example: minikube addons enable dashboard). For a list of available addons use: minikube addons list ",
	Long:  "Enables the addon w/ADDON_NAME within minikube (example: minikube addons enable dashboard). For a list of available addons use: minikube addons list ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintln(os.Stderr, "usage: minikube addons enable ADDON_NAME")
			os.Exit(1)
		}

		addon := args[0]
		err := Set(addon, "true")
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
		} else {
			fmt.Fprintln(os.Stdout, fmt.Sprintf("%s was successfully enabled", addon))
		}
	},
}

func init() {
	AddonsCmd.AddCommand(addonsEnableCmd)
}
