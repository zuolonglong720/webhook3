package controllers

import (	
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
	"github.com/Shopify/sarama"
    "time"
    "log"
    "fmt"
    "os"
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
   // fmt.Println(json)
	//fmt.Println(gjson.Get(json, "ref"))
    //从前端获得数据，并解包放入resp中
  
//	data := this.Ctx.Input.RequestBody
//	fmt.Println(&resp)
//	fmt.Println(string(resp["ref"]))
	//fmt.Println()

	fmt.Println("start kafka send")
	syncProducer(Address,json)

	this.Ctx.WriteString(json)
	fmt.Println("over ---")
}


var Address = []string{"47.97.248.41:9092"}


//同步消息模式
func syncProducer(address []string,json string)  {
    fmt.Printf("start")
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
     config.Producer.Timeout = 5 * time.Second
    p, err := sarama.NewSyncProducer(address, config)
    if err != nil {
        log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
        return
    }
    defer p.Close()
    topic := "my-test2"
   // srcValue := "sync: this is a zuolonglong. index=%d"
    //for i:=0; i<10; i++ {
        value := json
        msg := &sarama.ProducerMessage{
            Topic:topic,
            Value:sarama.ByteEncoder(value),
        }
        part, offset, err := p.SendMessage(msg)
        if err != nil {
            log.Printf("send message(%s) err=%s \n", value, err)
        }else {
            fmt.Fprintf(os.Stdout, value + "发送成功，partition=%d, offset=%d \n", part, offset)
        }
        time.Sleep(2*time.Second)
    //}
}