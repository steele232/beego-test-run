package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "worldclassmachinelearning.com"
	c.Data["Email"] = "jonathan.steele232@gmail.com"
	c.TplName = "index.tpl"
}
