package main

import (
	"fmt"
	apiblog "gindemo/api/blog"
	"gindemo/app/blog"
	"gindemo/app/shop"
	"gindemo/routers"
)

func main() {
	/*1*/
	// r := routers.SetupRouter()
	// if err := r.Run(":9000"); err != nil {
	// 	fmt.Printf("startup service failed, err:%v\n", err)
	// }

	/*2*/
	// r := gin.Default()
	// routers.LoadOrder(r)
	// routers.LoadUser(r)
	// if err := r.Run(); err != nil {
	// 	fmt.Printf("startup server failed,err:%#v\n", err)
	// }

	// 加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers, apiblog.Routers)
	//routers.Include(shop.Routers, blog1.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
