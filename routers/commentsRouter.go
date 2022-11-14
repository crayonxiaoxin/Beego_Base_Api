package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello_api/controllers:LoginController"] = append(beego.GlobalControllerRouter["hello_api/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:MediaController"] = append(beego.GlobalControllerRouter["hello_api/controllers:MediaController"],
        beego.ControllerComments{
            Method: "GetAllFiles",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:MediaController"] = append(beego.GlobalControllerRouter["hello_api/controllers:MediaController"],
        beego.ControllerComments{
            Method: "RemoveFile",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:MediaController"] = append(beego.GlobalControllerRouter["hello_api/controllers:MediaController"],
        beego.ControllerComments{
            Method: "GetMedia",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:RegisterController"] = append(beego.GlobalControllerRouter["hello_api/controllers:RegisterController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:UploadController"] = append(beego.GlobalControllerRouter["hello_api/controllers:UploadController"],
        beego.ControllerComments{
            Method: "UploadFile",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hello_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hello_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hello_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hello_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello_api/controllers:UserController"] = append(beego.GlobalControllerRouter["hello_api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
