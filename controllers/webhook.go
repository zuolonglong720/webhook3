package controllers

import (	
	"github.com/astaxie/beego"
	"fmt"
	"github.com/tidwall/gjson"
 )

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
        this.Ctx.WriteString("hello")
}
func (this *MainController)Post(){

	json := string(this.Ctx.Input.RequestBody)

    //一个文件一个结构体，为避免麻烦，就创建了个map
    fmt.Println(json)
	fmt.Println(gjson.Get(json, "ref"))
    //从前端获得数据，并解包放入resp中
  
//	data := this.Ctx.Input.RequestBody
//	fmt.Println(&resp)
//	fmt.Println(string(resp["ref"]))
	fmt.Println()

	this.Ctx.WriteString(json)
	fmt.Println("这是一个webhook的测试")
}

