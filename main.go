package main

import (
	_ "hello_api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 访问静态文件
	beego.BConfig.WebConfig.StaticDir["/uploads"] = "uploads"

	beego.Run()
}
