package main

import (
	"kubeyc/k8s-client-demo/utils"
	"log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main()  {
	client := utils.CreateClient(utils.CMD())
	// 1.9 起支持分块获取pod资源
	// Limit限制单次chunk传输的大小
	// 响应中包含一个Continue，继续获取数据在下次请求中携带这个值
	pods, err := client.CoreV1().Pods("").List(metav1.ListOptions{
		Limit: 2,
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		for _, podItem := range pods.Items {
			log.Println("client consume pod: ", podItem.Name)
		}


		if pods.Continue == "" {
			log.Println("client consume the full set of results.")
			break
		}

		log.Println("Continue = ", pods.Continue)

		pods, err = client.CoreV1().Pods("").List(metav1.ListOptions{
			Limit: 2,
			Continue: pods.Continue,
		})


		if err != nil {
			log.Fatal(err)
		}
	}
}