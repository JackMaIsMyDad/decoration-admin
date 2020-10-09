// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"decoration-admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/easy",
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/register", &controllers.UserController{}, "post:CreateUser"),
		beego.NSNamespace("/user",
			beego.NSBefore(controllers.TokenAuth),
			beego.NSRouter("/info", &controllers.UserController{}, "get:GetUserInfo"),
		),
	)
	beego.AddNamespace(ns)
}
