package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeyc/k8s-client-demo/utils"
	"log"
	kuberinformers "k8s.io/client-go/informers"
	"time"
)

func main()  {
	client := utils.CreateClient(utils.CMD())
	sharedINformerFactory := kuberinformers.NewSharedInformerFactoryWithOptions(client, time.Minute * 10)
	sharedINformerFactory.Core().V1().Pods().Informer()
	//pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, podItem := range pods.Items {
	//	fmt.Printf("Namespce: %s, Pod: %s\n", podItem.Namespace, podItem.Name)
	//}
	//
	//client.CoreV1().Pods("").List()
	watchI, err := client.CoreV1().Pods("").Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	result := <-watchI.ResultChan()
	fmt.Println(result.Object.(*v1.Pod).Name)
}