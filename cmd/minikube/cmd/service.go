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

package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"k8s.io/minikube/pkg/minikube/cluster"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/machine"
	"k8s.io/minikube/pkg/minikube/service"
)

const defaultServiceFormatTemplate = "http://{{.IP}}:{{.Port}}"

var (
	namespace          string
	https              bool
	serviceURLMode     bool
	serviceURLFormat   string
	serviceURLTemplate *template.Template
	wait               int
	interval           int
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service [flags] SERVICE",
	Short: "Gets the kubernetes URL(s) for the specified service in your local cluster",
	Long:  `Gets the kubernetes URL(s) for the specified service in your local cluster. In the case of multiple URLs they will be printed one at a time.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		t, err := template.New("serviceURL").Parse(serviceURLFormat)
		if err != nil {
			fmt.Fprintln(os.Stderr, "The value passed to --format is invalid:\n\n", err)
			os.Exit(1)
		}
		serviceURLTemplate = t

		RootCmd.PersistentPreRun(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) > 1 {
			errText := "Please specify a service name."
			fmt.Fprintln(os.Stderr, errText)
			os.Exit(1)
		}

		svc := args[0]
		api, err := machine.NewAPIClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting client: %s\n", err)
			os.Exit(1)
		}
		defer api.Close()

		cluster.EnsureMinikubeRunningOrExit(api, 1)
		err = service.WaitAndMaybeOpenService(api, namespace, svc,
			serviceURLTemplate, serviceURLMode, https, wait, interval)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening service: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	serviceCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "The service namespace")
	serviceCmd.Flags().BoolVar(&serviceURLMode, "url", false, "Display the kubernetes service URL in the CLI instead of opening it in the default browser")
	serviceCmd.Flags().BoolVar(&https, "https", false, "Open the service URL with https instead of http")
	serviceCmd.Flags().IntVar(&wait, "wait", constants.DefaultWait, "Amount of time to wait for a service in seconds")
	serviceCmd.Flags().IntVar(&interval, "interval", constants.DefaultWait, "The time interval for each check that wait performs in seconds")

	serviceCmd.PersistentFlags().StringVar(&serviceURLFormat, "format", defaultServiceFormatTemplate, "Format to output service URL in. This format will be applied to each url individually and they will be printed one at a time.")

	RootCmd.AddCommand(serviceCmd)
}
