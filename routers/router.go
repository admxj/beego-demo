package routers

import (
	"github.com/astaxie/beego"
	"imooc/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Include(&controllers.InfoController{})
	beego.Router("/", &controllers.GuessController{})

}
