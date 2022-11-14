package controllers

import (
	"hello_api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// ErrorController operations for Error
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	result := utils.Result{}
	result.ResultCode = utils.ERR_404
	c.Data["json"] = result
	c.ServeJSON()
}
