package controllers

import (
	"hello_api/models"

	beego "github.com/beego/beego/v2/server/web"
)

// 登入
type LoginController struct {
	beego.Controller
}

// @Title Login
// @Summary     登入
// @Description 登入
// @Param	username		query 	string	true		"用户名"
// @Param	password		query 	string	true		"密码"
// @Success 200 {object} utils.Result
// @Failure 403 user not exist
// @router / [post]
func (u *LoginController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	result := models.Login(&models.User{Username: username, Password: password})
	u.Data["json"] = result
	u.ServeJSON()
}
