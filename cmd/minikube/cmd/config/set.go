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

	pkgConfig "k8s.io/minikube/pkg/minikube/config"

	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set PROPERTY_NAME PROPERTY_VALUE",
	Short: "Sets an individual value in a minikube config file",
	Long: `Sets the PROPERTY_NAME config value to PROPERTY_VALUE
	These values can be overwritten by flags or environment variables at runtime.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, "usage: minikube config set PROPERTY_NAME PROPERTY_VALUE")
			os.Exit(1)
		}
		err := Set(args[0], args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(configSetCmd)
}

func Set(name string, value string) error {
	s, err := findSetting(name)
	if err != nil {
		return err
	}
	// Validate the new value
	err = run(name, value, s.validations)
	if err != nil {
		return err
	}

	// Set the value
	config, err := pkgConfig.ReadConfig()
	if err != nil {
		return err
	}
	err = s.set(config, name, value)
	if err != nil {
		return err
	}

	// Run any callbacks for this property
	err = run(name, value, s.callbacks)
	if err != nil {
		return err
	}

	// Write the value
	return WriteConfig(config)
}
