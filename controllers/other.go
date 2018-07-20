package controllers

import (
	"github.com/astaxie/beego"
)

type OtherController struct {
	beego.Controller
}

func (c *OtherController) Get() {
	c.Data["FirstName"] = "Jonathan"
	c.Data["LastName"] = "Steele"
	c.TplName = "other.tpl"
}
