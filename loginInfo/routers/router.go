// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"loginInfo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.Router("/regist",&controllers.RegistController{},"post:Insert")
	beego.Router("/login",&controllers.LoginController{},"post:Login")
	//beego.Router("/login1",&controllers.LoginController{},"get:Login1")
	//beego.Router("/token",&controllers.LoginController{},"get:Token")
	beego.Router("/orderadd",&controllers.OrderAddController{},"post:OrderAdd")
	beego.Router("/ordersearch",&controllers.OrderSearchController{},"post:OrderSearch")
	beego.Router("/orderupdate",&controllers.OrderSearchController{},"post:OrderUpdate")

	beego.Router("/orderdelete",&controllers.OrderSearchController{},"post:OrderDelete")
	//beego.Router("/one",&controllers.OrderSearchController{},"get:One")
	beego.AddNamespace(ns)
}
