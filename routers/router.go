package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
   	 beego.Router("/webhook", &controllers.MainController{})
}
