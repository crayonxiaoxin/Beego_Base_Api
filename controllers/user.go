package controllers

import (
	"hello_api/models"
	"hello_api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// 用户
type UserController struct {
	beego.Controller
}

// @Title GetAll
// @Description 获取所有用户
// @Param	token		header 	string	    true		"登入后返回的token"
// @Success 200 {object} utils.Result
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	var result = utils.Result{ResultCode: utils.SUCCESS, Data: users}
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title Get
// @Description 通过id获取用户
// @Param	token	header 	string	true		"登入后返回的token"
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} utils.Result
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid, err := u.GetInt(":uid", 0)
	var result = utils.Result{}
	if uid != 0 && err == nil {
		user := models.GetUser(uid)
		if user.Valid() {
			result.ResultCode = utils.SUCCESS
			result.Data = user
		} else {
			result.ResultCode = utils.ERR_USER_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title CreateUser
// @Description 添加用户
// @Param	token		header 	string	    true		"登入后返回的token"
// @Param	username		query 	string	true		"用户名"
// @Param	password		query 	string	true		"密码"
// @Success 200 {object} utils.Result
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	username := u.GetString("username", "")
	password := u.GetString("password", "")
	result := models.AddUser(&models.User{Username: username, Password: password})
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title Update
// @Description 更新用户
// @Param	token		header 	string	true		"登入后返回的token"
// @Param	uid			path 	string	true		"The uid you want to update"
// @Param	username	query 	string	true		"用户名"
// @Param	password	query 	string	false		"密码"
// @Success 200 {object} utils.Result
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid, err := u.GetInt(":uid", 0)
	if err != nil {
		uid = 0
	}
	username := u.GetString("username", "")
	password := u.GetString("password", "")
	user := models.User{Username: username, Password: password}
	user.ID = uint(uid)
	result := models.UpdateUser(&user)
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title Delete
// @Description 删除用户
// @Param	token	header 	string	true		"登入后返回的token"
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {object} utils.Result
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid, err := u.GetInt(":uid", 0)
	if err != nil {
		uid = 0
	}
	result := models.DeleteUser(uid)
	u.Data["json"] = result
	u.ServeJSON()
}
