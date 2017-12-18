package routers

import (
	"github.com/astaxie/beego"
	"github.com/pkufergus/rankserver/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/menu", &controllers.MenuController{})
}
