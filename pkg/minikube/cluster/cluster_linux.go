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

package cluster

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/docker/machine/libmachine/drivers"
	"k8s.io/minikube/pkg/drivers/none"
	cfg "k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/minikube/constants"
)

type kvmDriver struct {
	*drivers.BaseDriver

	Memory         int
	DiskSize       int
	CPU            int
	Network        string
	PrivateNetwork string
	ISO            string
	Boot2DockerURL string
	DiskPath       string
	CacheMode      string
	IOMode         string
}

func createKVMHost(config MachineConfig) *kvmDriver {
	return &kvmDriver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: cfg.GetMachineName(),
			StorePath:   constants.GetMinipath(),
			SSHUser:     "docker",
		},
		Memory:         config.Memory,
		CPU:            config.CPUs,
		Network:        config.KvmNetwork,
		PrivateNetwork: "docker-machines",
		Boot2DockerURL: config.Downloader.GetISOFileURI(config.MinikubeISO),
		DiskSize:       config.DiskSize,
		DiskPath:       filepath.Join(constants.GetMinipath(), "machines", cfg.GetMachineName(), fmt.Sprintf("%s.rawdisk", cfg.GetMachineName())),
		ISO:            filepath.Join(constants.GetMinipath(), "machines", cfg.GetMachineName(), "boot2docker.iso"),
		CacheMode:      "default",
		IOMode:         "threads",
	}
}

func createKVM2Host(config MachineConfig) *kvmDriver {
	return &kvmDriver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: cfg.GetMachineName(),
			StorePath:   constants.GetMinipath(),
			SSHUser:     "docker",
		},
		Memory:         config.Memory,
		CPU:            config.CPUs,
		Network:        config.KvmNetwork,
		PrivateNetwork: "minikube-net",
		Boot2DockerURL: config.Downloader.GetISOFileURI(config.MinikubeISO),
		DiskSize:       config.DiskSize,
		DiskPath:       filepath.Join(constants.GetMinipath(), "machines", cfg.GetMachineName(), fmt.Sprintf("%s.rawdisk", cfg.GetMachineName())),
		ISO:            filepath.Join(constants.GetMinipath(), "machines", cfg.GetMachineName(), "boot2docker.iso"),
		CacheMode:      "default",
		IOMode:         "threads",
	}
}

func detectVBoxManageCmd() string {
	cmd := "VBoxManage"
	if path, err := exec.LookPath(cmd); err == nil {
		return path
	}
	return cmd
}

func createNoneHost(config MachineConfig) *none.Driver {
	return &none.Driver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: cfg.GetMachineName(),
			StorePath:   constants.GetMinipath(),
		},
	}
}
