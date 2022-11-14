package controllers

import (
	"hello_api/models"
	"hello_api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// 文件
type MediaController struct {
	beego.Controller
}

// @Title       Delete
// @Summary     删除文件
// @Description 删除文件
// @Param       token header string true "token"
// @Param       id    path   int    true "文件ID"
// @Tags        media
// @Success     200 {object} utils.Result
// @router      /:id [delete]
func (c *MediaController) RemoveFile() {
	mid, err := c.GetInt(":id", 0)
	if err != nil {
		mid = 0
	}
	result := models.DeleteMedia(int(mid))
	c.Data["json"] = result
	c.ServeJSON()
}

// @Title       GetAll
// @Summary     获取所有媒体
// @Description 获取所有媒体
// @Param       token header string true  "登入后返回的token"
// @Param       page  query  int    false "页码"
// @Param       size  query  int    false "每页数量"
// @Tags        media
// @Success     200 {object} utils.Result
// @router      / [get]
func (c *MediaController) GetAllFiles() {
	page, err := c.GetInt("page", 0)
	if err != nil {
		page = 0
	}
	size, err := c.GetInt("size", 0)
	if err != nil {
		size = 0
	}
	users, count := models.GetAllMedia(page, size)
	data := make(map[string]interface{})
	data["list"] = users
	data["count"] = count
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: data}
	c.Data["json"] = result
	c.ServeJSON()
}

// @Title       Get
// @Summary     通过id获取文件
// @Description 通过id获取文件
// @Param       token header string true "登入后返回的token"
// @Param       id    path   int    true "The key for staticblock"
// @Tags        media
// @Success     200 {object} utils.Result
// @router      /:id [get]
func (c *MediaController) GetMedia() {
	mid, err := c.GetInt(":id", 0)
	if err != nil {
		mid = 0
	}
	var result = utils.Result{}
	if err == nil && mid != 0 {
		user := models.GetMedia(int(mid))
		if user.Valid() {
			result.ResultCode = utils.SUCCESS
			result.Data = user
		} else {
			result.ResultCode = utils.ERR_UPLOAD_FILE_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	c.Data["json"] = result
	c.ServeJSON()
}
