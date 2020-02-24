package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	name := c.GetString("name")
	c.Data["Website"] = name
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//bee generate scaffold info -fields="id:int64,name:string,age:int" -driver=mysql -conn="root:@tcp(127.0.0.1:3306)/nginx"
