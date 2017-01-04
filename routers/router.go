package routers

import (
	"stratacompare/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/managers", &controllers.MainController{}, "get:AllManagers")
	beego.Router("/financials", &controllers.MainController{}, "get:Financials")
	beego.Router("/services", &controllers.MainController{}, "get:Services")
	beego.Router("/short", &controllers.MainController{}, "get:ShortCodes")
}
