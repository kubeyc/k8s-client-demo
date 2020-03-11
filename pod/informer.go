package main

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"kubeyc/k8s-client-demo/utils"
	"log"
	"time"
)

func main() {
	client := utils.CreateClient(utils.CMD())
	sharedInformerFactory := informers.NewSharedInformerFactory(client, time.Minute)

	stopCh := make(chan struct{})
	defer close(stopCh)
	podInformer := sharedInformerFactory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Printf("New Pod Added to Store: %s", obj.(v1.Object).GetName())
		},

		DeleteFunc: func(obj interface{}) {
			log.Printf("New Pod Deleted to Store: %s", obj.(v1.Object).GetName())
		},
	})

	podInformer.Run(stopCh)
}
