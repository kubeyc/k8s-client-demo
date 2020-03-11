package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeyc/k8s-client-demo/utils"
	"log"
)

func main() {
	client := utils.CreateClient(utils.CMD())

	fmt.Println("Get All Namespace Pods:")
	pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, podItem := range pods.Items {
		fmt.Printf("Pod (Namespace = %s): %s\n", podItem.Namespace, podItem.Name)
	}

	fmt.Println("Get Pods Match Labels (tier=control-plane): ")
	pods, err = client.CoreV1().Pods("").List(metav1.ListOptions{
		LabelSelector: "tier=control-plane",

	})
	if err != nil {
		log.Fatal(err)
	}
	for _, podItem := range pods.Items {
		fmt.Printf("Pod (Namespace = %s， Label Match: tier=control-plane): %s\n",
			podItem.Namespace, podItem.Name)
	}
}
