package main

import (
    "fmt"
    "time"

     "github.com/golang/glog"

     "k8s.io/api/core/v1"
     "k8s.io/apimachinery/pkg/fields"
     "k8s.io/client-go/kubernetes"
     "k8s.io/client-go/tools/cache"
     "k8s.io/client-go/tools/clientcmd"
)

func main() {
    config, err := clientcmd.BuildConfigFromFlags("", "")
    if err != nil {
        glog.Errorln(err)
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        glog.Errorln(err)
    }

    watchlist := cache.NewListWatchFromClient(
        clientset.CoreV1().RESTClient(),
        string(v1.ResourceServices),
        v1.NamespaceAll,
        fields.Everything(),
    )
    _, controller := cache.NewInformer(
        watchlist,
        &v1.Service{},
        0, //Duration is int64
        cache.ResourceEventHandlerFuncs{
            AddFunc: func(obj interface{}) {
                fmt.Printf("service added: %s \n", obj)
            },
            DeleteFunc: func(obj interface{}) {
                fmt.Printf("service deleted: %s \n", obj)
            },
            UpdateFunc: func(oldObj, newObj interface{}) {
                fmt.Printf("service changed \n")
            },
         },
     )
    stop := make(chan struct{})
    defer close(stop)
    go controller.Run(stop)
    for {
        time.Sleep(time.Second)
    }
}
