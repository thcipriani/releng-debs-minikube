/*
Copyright 2017 The Kubernetes Authors All rights reserved.

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
	"github.com/spf13/cobra"
	cmdConfig "k8s.io/minikube/cmd/minikube/cmd/config"
	"k8s.io/minikube/pkg/minikube/config"
	"k8s.io/minikube/pkg/minikube/constants"
	"k8s.io/minikube/pkg/minikube/machine"
	"os"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Add or delete an image from the local cache.",
	Long:  "Add or delete an image from the local cache.",
}

// addCacheCmd represents the cache add command
var addCacheCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an image to local cache.",
	Long:  "Add an image to local cache.",
	Run: func(cmd *cobra.Command, args []string) {
		// Cache and load images into docker daemon
		if err := machine.CacheAndLoadImages(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error caching and loading images: %s\n", err)
			os.Exit(1)
		}
		// Add images to config file
		if err := cmdConfig.AddToConfigMap(constants.Cache, args); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding cached images to config file: %s\n", err)
			os.Exit(1)
		}
	},
}

// deleteCacheCmd represents the cache delete command
var deleteCacheCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an image from the local cache.",
	Long:  "Delete an image from the local cache.",
	Run: func(cmd *cobra.Command, args []string) {
		// Delete images from config file
		if err := cmdConfig.DeleteFromConfigMap(constants.Cache, args); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting images from config file: %s\n", err)
			os.Exit(1)
		}
		// Delete images from cache/images directory
		if err := machine.DeleteFromImageCacheDir(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting images: %s\n", err)
			os.Exit(1)
		}
	},
}

// LoadCachedImagesInConfigFile loads the images currently in the config file (minikube start)
func LoadCachedImagesInConfigFile() error {
	configFile, err := config.ReadConfig()
	if err != nil {
		return err
	}
	if values, ok := configFile[constants.Cache]; ok {
		var images []string
		for key := range values.(map[string]interface{}) {
			images = append(images, key)
		}
		return machine.CacheAndLoadImages(images)
	}
	return nil
}

func init() {
	cacheCmd.AddCommand(addCacheCmd)
	cacheCmd.AddCommand(deleteCacheCmd)
	RootCmd.AddCommand(cacheCmd)
}
