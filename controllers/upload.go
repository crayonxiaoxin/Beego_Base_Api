package controllers

import (
	"hello_api/models"
	"hello_api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// 上传
type UploadController struct {
	beego.Controller
}

// @Title       Upload
// @Summary     上传
// @Description 上传
// @Param       token header   string true "token"
// @Param       file  formData file   true "文件"
// @Tags        media
// @Success     200 {object} utils.Result
// @router      / [post]
func (c *UploadController) UploadFile() {
	f, fh, err := c.GetFile("file")
	result := utils.Result{}
	if err != nil {
		result.ResultCode = utils.ERR_PARAMS
	} else {
		defer f.Close()
		result = models.UploadMedia(fh)
	}
	c.Data["json"] = result
	c.ServeJSON()
}
