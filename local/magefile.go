//go:build mage
// +build mage

// Contains targets to help setup a full kubeflow deployment
// Currently a MVP
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Make sure to run mage clean before changing variables to avoid clutter
const CLUSTER_NAME = "kfcluster"

func removeClusterOnInterrupt(c chan os.Signal) {
	<-c
	sh.Run("kind", "delete", "cluster", "--name", CLUSTER_NAME)
}

func Cluster() error {
	clusterName := fmt.Sprintf("kind-%s", CLUSTER_NAME)
	if _, err := sh.Output("kubectl", "cluster-info", "--context", clusterName); err == nil {
		fmt.Println("Cluster already exists...Skipping this step")
		return nil
	}

	// kind doesn't clean a cluster creation in progress on keyboard interrupts
	// Issue when a cluster already exists. Can cause accidental deletion
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// go removeClusterOnInterrupt(c)

	return sh.RunV("kind", "create", "cluster", "--name", CLUSTER_NAME)
}

// Checks the reason why the kubeflow setup is not up
func Status() error {
	return nil
}

// Checks whether the tools required to run this mage file are installed. Utility
func Prerequisites() error {
	return nil
}

// Clone the manifests repository
func Manifests() error {
	if _, err := os.Stat("./manifests"); os.IsNotExist(err) {
		return sh.RunV("git", "clone", "https://github.com/kubeflow/manifests")
	}
	fmt.Println("manifests folder already exists...Skipping this step.")
	return nil
}

// performs a full deployment of kubeflow
// Reference: https://github.com/kubeflow/manifests#install-with-a-single-command
// `while ! kustomize build example | kubectl apply -f -; do echo "Retrying to apply resources"; sleep 10; done`
func Full() error {
	mg.Deps(Cluster, Manifests)
	os.Chdir("manifests")
	defer os.Chdir("..")

	var err error
	for {
		err = sh.RunV("kubectl", "apply", "-k", "example")
		if err == nil {
			break
		}
		fmt.Println("Retrying to apply resources")
		time.Sleep(time.Second * 10)
	}
	return err
}

func Clean() error {
	return sh.RunV("kind", "delete", "cluster", "--name", CLUSTER_NAME)
}
