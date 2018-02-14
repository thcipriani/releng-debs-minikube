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
	"net"

	flag "github.com/spf13/pflag"

	"k8s.io/minikube/pkg/localkube"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/util"
)

func NewLocalkubeServer() *localkube.LocalkubeServer {
	_, defaultServiceCIDR, _ := net.ParseCIDR(util.DefaultServiceCIDR)
	return &localkube.LocalkubeServer{
		Containerized:            false,
		DNSDomain:                util.DefaultDNSDomain,
		LocalkubeDirectory:       util.DefaultLocalkubeDirectory,
		APIServerAddress:         net.ParseIP("0.0.0.0"),
		APIServerPort:            util.APIServerPort,
		APIServerInsecureAddress: net.ParseIP("127.0.0.1"),
		APIServerInsecurePort:    0,
		APIServerName:            constants.APIServerName,
		ShouldGenerateKubeconfig: false,
		ShouldGenerateCerts:      true,
		ShowVersion:              false,
		ServiceClusterIPRange:    *defaultServiceCIDR,
		RuntimeConfig:            map[string]string{"api/all": "true"},
		ExtraConfig:              util.ExtraOptionSlice{},
	}
}

// AddFlags adds flags for a specific LocalkubeServer
func AddFlags(s *localkube.LocalkubeServer) {
	flag.BoolVar(&s.Containerized, "containerized", s.Containerized, "If kubelet should run in containerized mode")
	flag.BoolVar(&s.EnableDNS, "enable-dns", s.EnableDNS, "DEPRECATED: Please run kube-dns as a cluster addon")
	flag.StringVar(&s.DNSDomain, "dns-domain", s.DNSDomain, "The cluster dns domain")
	flag.StringVar(&s.LocalkubeDirectory, "localkube-directory", s.LocalkubeDirectory, "The directory localkube will store files in")
	flag.IPNetVar(&s.ServiceClusterIPRange, "service-cluster-ip-range", s.ServiceClusterIPRange, "The service-cluster-ip-range for the apiserver")
	flag.IPVar(&s.APIServerAddress, "apiserver-address", s.APIServerAddress, "The address the apiserver will listen securely on")
	flag.IntVar(&s.APIServerPort, "apiserver-port", s.APIServerPort, "The port the apiserver will listen securely on")
	flag.IPVar(&s.APIServerInsecureAddress, "apiserver-insecure-address", s.APIServerInsecureAddress, "The address the apiserver will listen insecurely on")
	flag.IntVar(&s.APIServerInsecurePort, "apiserver-insecure-port", s.APIServerInsecurePort, "The port the apiserver will listen insecurely on")
	flag.StringVar(&s.APIServerName, "apiserver-name", s.APIServerName, "The apiserver name which is used in the generated certificate for localkube/kubernetes.  This can be used if you want to make the API server available from outside the machine")

	flag.BoolVar(&s.ShouldGenerateKubeconfig, "generate-kubeconfig", s.ShouldGenerateKubeconfig, "If localkube should generate its own kubeconfig")
	flag.BoolVar(&s.ShouldGenerateCerts, "generate-certs", s.ShouldGenerateCerts, "If localkube should generate it's own certificates")
	flag.BoolVar(&s.ShowVersion, "show-version", s.ShowVersion, "If localkube should just print the version and exit.")
	flag.BoolVar(&s.ShowHostIP, "host-ip", s.ShowHostIP, "If localkube should just print the host IP and exit.")
	flag.Var(&s.RuntimeConfig, "runtime-config", "A set of key=value pairs that describe runtime configuration that may be passed to apiserver. apis/<groupVersion> key can be used to turn on/off specific api versions. apis/<groupVersion>/<resource> can be used to turn on/off specific resources. api/all and api/legacy are special keys to control all and legacy api versions respectively.")
	flag.IPVar(&s.NodeIP, "node-ip", s.NodeIP, "IP address of the node. If set, kubelet will use this IP address for the node.")
	flag.StringVar(&s.ContainerRuntime, "container-runtime", "", "The container runtime to be used")
	flag.StringVar(&s.RemoteRuntimeEndpoint, "remote-runtime-endpoint", "", "The container runtime endpoint (CRI) to be used (if this is set, then --container-runtime is forced as 'remote')")
	flag.StringVar(&s.RemoteImageEndpoint, "remote-image-endpoint", "", "The container image endpoint (CRI) to be used (if this is set, then --container-runtime is forced as 'remote')")
	flag.StringVar(&s.NetworkPlugin, "network-plugin", "", "The name of the network plugin")
	flag.StringVar(&s.FeatureGates, "feature-gates", "", "A set of key=value pairs that describe feature gates for alpha/experimental features.")
	flag.Var(&s.ExtraConfig, "extra-config", "A set of key=value pairs that describe configuration that may be passed to different components. The key should be '.' separated, and the first part before the dot is the component to apply the configuration to.")

	// These two come from vendor/ packages that use flags. We should hide them
	flag.CommandLine.MarkHidden("google-json-key")
	flag.CommandLine.MarkHidden("log-flush-frequency")

	// Parse them
	flag.Parse()
}
