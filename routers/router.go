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
		//beego.NSNamespace("/object",
		//	beego.NSInclude(
		//		&controllers.ObjectController{},
		//	),
		//),
		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:CreateUser"),
			beego.NSRouter("info", &controllers.UserController{}, "get:GetUserInfo"),
			//beego.NSAutoRouter(&controllers.UserController{}),
			//beego.NSInclude(&controllers.UserController{}),
		),
	)
	beego.AddNamespace(ns)
}
