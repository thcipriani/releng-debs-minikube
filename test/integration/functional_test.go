// +build integration

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

package integration

import (
	"strings"
	"testing"
)

func TestFunctional(t *testing.T) {
	minikubeRunner := NewMinikubeRunner(t)
	minikubeRunner.EnsureRunning()
	// This one is not parallel, and ensures the cluster comes up
	// before we run any other tests.
	t.Run("Status", testClusterStatus)

	t.Run("DNS", testClusterDNS)
	t.Run("Logs", testClusterLogs)
	t.Run("Addons", testAddons)
	t.Run("Dashboard", testDashboard)
	t.Run("ServicesList", testServicesList)
	t.Run("Provisioning", testProvisioning)

	if !strings.Contains(minikubeRunner.StartArgs, "--vm-driver=none") {
		t.Run("EnvVars", testClusterEnv)
		t.Run("SSH", testClusterSSH)
		t.Run("IngressController", testIngressController)
		// t.Run("Mounting", testMounting)
	}
}
