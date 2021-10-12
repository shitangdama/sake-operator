package main

// import (
// 	"context"
// 	"fmt"

// 	rest "k8s.io/client-go/rest"
// 	"k8s.io/client-go/tools/cache"
// )

// // InitController init ControllerManager
// func InitController(ctx context.Context, config *rest.Config) error {
// 	deployFactory := kubeinformers.NewSharedInformerFactory(deployClient, defaultResync)
// 	Controller.DeploymentInformer = deployFactory.Apps().V1().Deployments()

// 	// context 更换
// 	go deployFactory.Start(ctx.Done())
// 	// go istioFactory.Start(ctx.Done())

// 	Controller.DeploymentHasSynced = Controller.DeploymentInformer.Informer().HasSynced

// 	if !cache.WaitForCacheSync(ctx.Done(), Controller.DeploymentInformer.Informer().HasSynced, Controller.VirtualServiceInformer.Informer().HasSynced) {
// 		common.Log.Error(fmt.Errorf("Timed out waiting for caches to sync"))
// 		return nil
// 	}

// 	Controller.DeploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
// 		AddFunc:    onAddDeployment,
// 		UpdateFunc: onUpdateDeployment,
// 		DeleteFunc: onDeleteDeployment,
// 	})
// }

// // 放入queue触发对应的
