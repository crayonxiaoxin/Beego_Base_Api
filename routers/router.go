// @APIVersion v1
// @Title Test API
// @Description Beego API 基础工程
package routers

import (
	"hello_api/controllers"
	"hello_api/utils"

	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	// 直接注册不能使用注解路由，需要使用 namespace
	// beego.Router("/authorize", &controllers.TokenController{})

	// 设置过滤器（token中间件）
	NeedToken("/v1/user/?:id")
	// NeedToken("/v1/user/*")

	// 404
	beego.ErrorController(&controllers.ErrorController{})
}

// 验证 token
func ValifyToken(ctx *beecontext.Context) {
	tokenString := ctx.Input.Header("token")
	rc, _ := utils.ParseToken(tokenString)
	if !rc.Success() {
		ctx.Output.JSON(rc, true, false)
	}
}

func NeedToken(pattern string) {
	// 设置过滤器（token中间件）
	beego.InsertFilter(pattern, beego.BeforeRouter, ValifyToken)
}
