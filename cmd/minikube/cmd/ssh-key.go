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

	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/minikube/constants"
)

// sshKeyCmd represents the sshKey command
var sshKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Retrieve the ssh identity key path of the specified cluster",
	Long:  "Retrieve the ssh identity key path of the specified cluster.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(filepath.Join(constants.GetMinipath(), "machines", config.GetMachineName(), "id_rsa"))
	},
}

func init() {
	RootCmd.AddCommand(sshKeyCmd)
}
