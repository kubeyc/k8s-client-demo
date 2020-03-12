package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"kubeyc/k8s-client-demo/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatal(err)
	}
	restConfig.AcceptContentTypes = "application/json;as=Table;g=meta.k8s.io;v=v1beta1"
	client := utils.CreateRESTClient(restConfig)
	pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, podItem := range pods.Items {
		log.Println(podItem.Name)
	}
}
