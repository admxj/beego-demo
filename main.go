package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "immooc/routers"
)

func main() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/nginx?charset=utf8")
	beego.Run()
}
