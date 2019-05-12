package routers

import (
	"webhook3/controllers"
	"github.com/astaxie/beego"
)

func init() {
   	 beego.Router("/webhook", &controllers.MainController{})
}
