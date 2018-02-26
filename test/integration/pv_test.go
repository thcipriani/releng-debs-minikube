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
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/labels"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/storage"
	commonutil "k8s.io/minikube/pkg/util"
	"k8s.io/minikube/test/integration/util"
)

var (
	pvcName = "testpvc"
	pvcCmd  = []string{"get", "pvc", pvcName}
)

func testProvisioning(t *testing.T) {
	t.Parallel()
	kubectlRunner := util.NewKubectlRunner(t)

	defer func() {
		kubectlRunner.RunCommand([]string{"delete", "pvc", pvcName})
	}()

	// We have to make sure the addon-manager has created the StorageClass before creating
	// a claim. Otherwise it will never get bound.

	checkStorageClass := func() error {
		scl := storage.StorageClassList{}
		kubectlRunner.RunCommandParseOutput([]string{"get", "storageclass"}, &scl)
		if len(scl.Items) > 0 {
			return nil
		}
		return fmt.Errorf("No default StorageClass yet.")
	}

	if err := util.Retry(t, checkStorageClass, 5*time.Second, 20); err != nil {
		t.Fatalf("No default storage class: %s", err)
	}

	// Check that the storage provisioner pod is running

	checkPodRunning := func() error {
		client, err := commonutil.GetClient()
		if err != nil {
			return errors.Wrap(err, "getting kubernetes client")
		}
		selector := labels.SelectorFromSet(labels.Set(map[string]string{"integration-test": "storage-provisioner"}))

		if err := commonutil.WaitForPodsWithLabelRunning(client, "kube-system", selector); err != nil {
			return err
		}
		return nil
	}

	if err := checkPodRunning(); err != nil {
		t.Fatal("Check storage-provisioner pod running failed with error: ", err)
	}

	// Now create the PVC
	pvcPath := filepath.Join(*testdataDir, "pvc.yaml")
	if _, err := kubectlRunner.RunCommand([]string{"create", "-f", pvcPath}); err != nil {
		t.Fatalf("Error creating pvc")
	}

	// And check that it gets bound to a PV.
	checkStorage := func() error {
		pvc := api.PersistentVolumeClaim{}
		if err := kubectlRunner.RunCommandParseOutput(pvcCmd, &pvc); err != nil {
			return err
		}
		// The test passes if the volume claim gets bound.
		if pvc.Status.Phase == "Bound" {
			return nil
		}
		return fmt.Errorf("PV not attached to PVC: %v", pvc)
	}

	if err := util.Retry(t, checkStorage, 2*time.Second, 5); err != nil {
		t.Fatal("PV Creation failed with error:", err)
	}

}
