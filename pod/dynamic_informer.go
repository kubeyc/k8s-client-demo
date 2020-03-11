package main

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"kubeyc/k8s-client-demo/utils"
	"log"
	"os"
	"os/signal"
)

func main() {
	restConfig, err := clientcmd.BuildConfigFromFlags("", utils.CMD())
	if err != nil {
		log.Fatal(err)
	}

	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		log.Fatal(err)
	}

	dynamicSharedInformerFactory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicClient, 0, "default", nil)
	gvr, _ := schema.ParseResourceArg("deployments.v1.core")
	resourceInformer := dynamicSharedInformerFactory.ForResource(*gvr)
	sharedIndexInformer := resourceInformer.Informer()
	sharedIndexInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Println("received add event!")
		},
		UpdateFunc: func(oldObj, obj interface{}) {
			log.Println("received update event!")
		},
		DeleteFunc: func(obj interface{}) {
			log.Println("received update event!")
		},
	})

	stopCh := make(chan struct{})
	sharedIndexInformer.Run(stopCh)

	sigCh := make(chan os.Signal, 0)
	signal.Notify(sigCh, os.Kill, os.Interrupt)

	<-sigCh
	close(stopCh)
}
