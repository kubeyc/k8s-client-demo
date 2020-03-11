package main

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"os/signal"
)

type clusterController struct {
	clientset *kubernetes.Clientset
}

func (c *clusterController) Run() {
	// 使用Factory方式创建一个Shared Informer来catch事件
	factory := informers.NewSharedInformerFactory(c.clientset, 0)
	// 定义需要catch对资源， 这里是Pod
	informer := factory.Core().V1().Pods().Informer()

	// 用于通知informer loop关闭
	stopper := make(chan struct{})
	defer close(stopper)

	// 定义处理crash对函数
	defer runtime.HandleCrash()

	// 添加监听事件对handler function
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {

		},
		DeleteFunc: func(obj interface{}) {

		},
		UpdateFunc: func(oldObj, newObj interface{}) {

		},
	})

	informer.Run(stopper)
}

func main() {
	// 首先需要初始化一个客户端
	// 下面对代码将打开kube-config文件并初始化客户端cmd(也就是打开到kube-apiserver的连接)
	config, err := clientcmd.BuildConfigFromFlags("", "~/.kube/config")
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err.Error())
	}

	c := clusterController{
		clientset: clientset,
	}

	go c.Run()

	sigCh := make(chan os.Signal, 0)
	signal.Notify(sigCh, os.Kill, os.Interrupt)

	<-sigCh
}
