package routers

import (
	"github.com/astaxie/beego"
	"github.com/steele232/testbeego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/FullName", &controllers.OtherController{})
}
