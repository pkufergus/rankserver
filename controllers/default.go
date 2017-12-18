package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type MenuController struct {
	beego.Controller
}
type MenuResp struct {
	Id       uint64
	Name     string
	FatherId uint64
}

func (c *MenuController) Get() {
	m := MenuResp{}
	m.Id = 1
	m.Name = "car"
	m.FatherId = 0
	c.Data["json"] = &m
	c.ServeJSON()
}
