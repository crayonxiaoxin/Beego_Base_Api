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
// @Summary     注册
// @Description 注册
// @Param	username		query 	string	true		"用户名"
// @Param	password		query 	string	true		"密码"
// @Success 200 {object} utils.Result
// @Failure 403 user not exist
// @router / [post]
func (c *RegisterController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")
	result := models.Register(&models.User{Username: username, Password: password})
	c.Data["json"] = result
	c.ServeJSON()
}
