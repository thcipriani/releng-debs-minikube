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

package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/util"
)

type Addon struct {
	Assets    []*BinDataAsset
	enabled   bool
	addonName string
}

func NewAddon(assets []*BinDataAsset, enabled bool, addonName string) *Addon {
	a := &Addon{
		Assets:    assets,
		enabled:   enabled,
		addonName: addonName,
	}
	return a
}

func (a *Addon) IsEnabled() (bool, error) {
	addonStatusText, err := config.Get(a.addonName)
	if err == nil {
		addonStatus, err := strconv.ParseBool(addonStatusText)
		if err != nil {
			return false, err
		}
		return addonStatus, nil
	}
	return a.enabled, nil
}

var Addons = map[string]*Addon{
	"addon-manager": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/addon-manager.yaml",
			"/etc/kubernetes/manifests/",
			"addon-manager.yaml",
			"0640"),
	}, true, "addon-manager"),
	"dashboard": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/dashboard/dashboard-dp.yaml",
			constants.AddonsPath,
			"dashboard-dp.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/dashboard/dashboard-svc.yaml",
			constants.AddonsPath,
			"dashboard-svc.yaml",
			"0640"),
	}, true, "dashboard"),
	"default-storageclass": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/storageclass/storageclass.yaml",
			constants.AddonsPath,
			"storageclass.yaml",
			"0640"),
	}, true, "default-storageclass"),
	"storage-provisioner": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/storage-provisioner/storage-provisioner.yaml",
			constants.AddonsPath,
			"storage-provisioner.yaml",
			"0640"),
	}, true, "storage-provisioner"),
	"coredns": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/coredns/coreDNS-controller.yaml",
			constants.AddonsPath,
			"coreDNS-controller.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/coredns/coreDNS-configmap.yaml",
			constants.AddonsPath,
			"coreDNS-configmap.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/coredns/coreDNS-svc.yaml",
			constants.AddonsPath,
			"coreDNS-svc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/coredns/coreDNS-crbinding.yaml",
			constants.AddonsPath,
			"coreDNS-crbinding.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/coredns/coreDNS-sa.yaml",
			constants.AddonsPath,
			"coreDNS-sa.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/coredns/coreDNS-clusterrole.yaml",
			constants.AddonsPath,
			"coreDNS-clusterrole.yaml",
			"0640"),
	}, false, "coredns"),
	"kube-dns": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/kube-dns/kube-dns-controller.yaml",
			constants.AddonsPath,
			"kube-dns-controller.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/kube-dns/kube-dns-cm.yaml",
			constants.AddonsPath,
			"kube-dns-cm.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/kube-dns/kube-dns-svc.yaml",
			constants.AddonsPath,
			"kube-dns-svc.yaml",
			"0640"),
	}, true, "kube-dns"),
	"heapster": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/heapster/influx-grafana-rc.yaml",
			constants.AddonsPath,
			"influxGrafana-rc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/heapster/grafana-svc.yaml",
			constants.AddonsPath,
			"grafana-svc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/heapster/influxdb-svc.yaml",
			constants.AddonsPath,
			"influxdb-svc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/heapster/heapster-rc.yaml",
			constants.AddonsPath,
			"heapster-rc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/heapster/heapster-svc.yaml",
			constants.AddonsPath,
			"heapster-svc.yaml",
			"0640"),
	}, false, "heapster"),
	"efk": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/efk/elasticsearch-rc.yaml",
			constants.AddonsPath,
			"elasticsearch-rc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/efk/elasticsearch-svc.yaml",
			constants.AddonsPath,
			"elasticsearch-svc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/efk/fluentd-es-rc.yaml",
			constants.AddonsPath,
			"fluentd-es-rc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/efk/fluentd-es-configmap.yaml",
			constants.AddonsPath,
			"fluentd-es-configmap.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/efk/kibana-rc.yaml",
			constants.AddonsPath,
			"kibana-rc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/efk/kibana-svc.yaml",
			constants.AddonsPath,
			"kibana-svc.yaml",
			"0640"),
	}, false, "efk"),
	"ingress": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/ingress/ingress-configmap.yaml",
			constants.AddonsPath,
			"ingress-configmap.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/ingress/ingress-rc.yaml",
			constants.AddonsPath,
			"ingress-rc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/ingress/ingress-svc.yaml",
			constants.AddonsPath,
			"ingress-svc.yaml",
			"0640"),
	}, false, "ingress"),
	"registry": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/registry/registry-rc.yaml",
			constants.AddonsPath,
			"registry-rc.yaml",
			"0640"),
		NewBinDataAsset(
			"deploy/addons/registry/registry-svc.yaml",
			constants.AddonsPath,
			"registry-svc.yaml",
			"0640"),
	}, false, "registry"),
	"registry-creds": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/registry-creds/registry-creds-rc.yaml",
			constants.AddonsPath,
			"registry-creds-rc.yaml",
			"0640"),
	}, false, "registry-creds"),
	"freshpod": NewAddon([]*BinDataAsset{
		NewBinDataAsset(
			"deploy/addons/freshpod/freshpod-rc.yaml",
			constants.AddonsPath,
			"freshpod-rc.yaml",
			"0640"),
	}, false, "freshpod"),
}

func AddMinikubeDirAssets(assets *[]CopyableFile) error {
	if err := addMinikubeDirToAssets(constants.MakeMiniPath("addons"), constants.AddonsPath, assets); err != nil {
		return errors.Wrap(err, "adding addons folder to assets")
	}
	if err := addMinikubeDirToAssets(constants.MakeMiniPath("files"), "", assets); err != nil {
		return errors.Wrap(err, "adding files rootfs to assets")
	}

	return nil
}

// AddMinikubeDirToAssets adds all the files in the basedir argument to the list
// of files to be copied to the vm.  If vmpath is left blank, the files will be
// transferred to the location according to their relative minikube folder path.
func addMinikubeDirToAssets(basedir, vmpath string, assets *[]CopyableFile) error {
	err := filepath.Walk(basedir, func(hostpath string, info os.FileInfo, err error) error {
		isDir, err := util.IsDirectory(hostpath)
		if err != nil {
			return errors.Wrapf(err, "checking if %s is directory", hostpath)
		}
		if !isDir {
			if vmpath == "" {
				rPath, err := filepath.Rel(basedir, hostpath)
				if err != nil {
					return errors.Wrap(err, "generating relative path")
				}
				rPath = filepath.Dir(rPath)
				vmpath = filepath.Join("/", rPath)
			}
			permString := fmt.Sprintf("%o", info.Mode().Perm())
			// The conversion will strip the leading 0 if present, so add it back
			// if we need to.
			if len(permString) == 3 {
				permString = fmt.Sprintf("0%s", permString)
			}

			f, err := NewFileAsset(hostpath, vmpath, filepath.Base(hostpath), permString)
			if err != nil {
				return errors.Wrapf(err, "creating file asset for %s", hostpath)
			}
			*assets = append(*assets, f)
		}

		return nil
	})
	if err != nil {
		return errors.Wrap(err, "walking filepath")
	}
	return nil
}
