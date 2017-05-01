package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}
type GetNameController struct {
	beego.Controller
}

type RuleHomeController struct {
	beego.Controller
}
type ListHomeController struct {
	beego.Controller
}

func (c *MainController) Get() {
	t := c.Ctx.Input.Header("Accept-Language")
	logs.Info("----------------- %s cat is %v years old", t, 3)
	v := c.GetSession("asta")
	if v == nil {
		c.SetSession("asta", int(1))
		c.Data["num"] = 0
	} else {
		c.SetSession("asta", v.(int)+1)
		c.Data["num"] = v.(int)
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}

func (c *RuleHomeController) Get() {
	c.TplName = "rule.html"
}

func (c *ListHomeController) Get() {
	c.TplName = "list.html"
}

func (this *GetNameController) Get() {
	myjson := `{"status":"200"}`
	this.Data["json"] = &myjson
	this.ServeJSON()

	//myjson :=`{"status":"200"}`
	//this.Data["xml"] =&myjson
	//this.ServeXML()
}
