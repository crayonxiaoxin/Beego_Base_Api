package controllers

import (
	"hello_api/models"

	beego "github.com/beego/beego/v2/server/web"
)

// 注册
type RegisterController struct {
	beego.Controller
}

// @Title Register
// @Description 注册
// @Param	username		query 	string	true		"用户名"
// @Param	password		query 	string	true		"密码"
// @Success 200 {object} utils.Result
// @Failure 403 user not exist
// @router / [post]
func (u *RegisterController) Register() {
	username := u.GetString("username")
	password := u.GetString("password")
	result := models.Register(&models.User{Username: username, Password: password})
	u.Data["json"] = result
	u.ServeJSON()
}
