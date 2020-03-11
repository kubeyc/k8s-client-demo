package main

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"kubeyc/k8s-client-demo/utils"
	"log"
)

func main()  {
	client := utils.CreateClient(utils.CMD())
	pods, err := client.CoreV1().Pods("demo").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	watchInterface, err := client.CoreV1().Pods("demo").Watch(metav1.ListOptions{
		Watch:               true,
		ResourceVersion:     pods.ResourceVersion,
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		event := <-watchInterface.ResultChan()
		pod := event.Object.(*v1.Pod)
		switch event.Type {
		case watch.Added:
			log.Printf("%s pod created", pod.Name)
		case watch.Modified:
			log.Printf("%s pod modified", pod.Name)
		case watch.Deleted:
			log.Printf("%s pod deleted", pod.Name)
		case watch.Error:
		case watch.Bookmark:
			log.Printf("bookmark, rv = %s", pod.ResourceVersion)
		}
	}
}