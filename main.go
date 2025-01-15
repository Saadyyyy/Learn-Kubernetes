package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Path ke kubeconfig (default di ~/.kube/config)
	kubeconfig := flag.String("kubeconfig", filepath.Join(homeDir(), ".kube", "config"), "Path to the kubeconfig file")
	flag.Parse()

	// Buat konfigurasi klien dari kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(fmt.Errorf("failed to load kubeconfig: %v", err))
	}

	// Buat klien Kubernetes
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Errorf("failed to create Kubernetes client: %v", err))
	}

	// Dapatkan daftar pod di namespace default
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to list pods: %v", err))
	}

	// Cetak nama-nama pod
	fmt.Println("Pods in the 'default' namespace:")
	for _, pod := range pods.Items {
		fmt.Printf("- %s (Status: %s)\n", pod.Name, pod.Status.Phase)
	}
}

// Fungsi untuk mendapatkan direktori home pengguna
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // Windows
}
