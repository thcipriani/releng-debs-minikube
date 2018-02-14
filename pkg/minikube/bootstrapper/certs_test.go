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

package bootstrapper

import (
	"os"
	"path/filepath"
	"testing"

	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/tests"
	"k8s.io/minikube/pkg/util"
)

func TestSetupCerts(t *testing.T) {
	tempDir := tests.MakeTempDir()
	defer os.RemoveAll(tempDir)

	f := NewFakeCommandRunner()
	k8s := KubernetesConfig{
		APIServerName: constants.APIServerName,
		DNSDomain:     constants.ClusterDNSDomain,
		ServiceCIDR:   util.DefaultServiceCIDR,
	}

	var filesToBeTransferred []string
	for _, cert := range certs {
		filesToBeTransferred = append(filesToBeTransferred, filepath.Join(constants.GetMinipath(), cert))
	}

	if err := SetupCerts(f, k8s); err != nil {
		t.Fatalf("Error starting cluster: %s", err)
	}
	for _, cert := range filesToBeTransferred {
		_, err := f.GetFileToContents(cert)
		if err != nil {
			t.Errorf("Cert not generated: %s", cert)
		}
	}
}
