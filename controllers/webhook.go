package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tidwall/gjson"
)

type Param struct {
	URL          string
	RepositoryId string
	Added        string
	Removed      string
	Modified     string
	Commits      string
}

var Address = []string{"47.97.248.41:9092"}
var _topic = "my-test2"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello")
}
func (this *MainController) Post() {

	json := string(this.Ctx.Input.RequestBody)

	//一个文件一个结构体，为避免麻烦，就创建了个map
	// fmt.Println(json)
	/*	 p := Param{}
	commits := gjson.Get(json, "commits")
	p.Commits = commits.String()

	restoryId := gjson.Get(json, "repository.id")
	p.RepositoryId = restoryId.String();*/

	var addedResult []string
	addeds := gjson.Get(json, "commits.#.added")
	for _, add := range addeds.Array() {
		fmt.Println(add)
		for _, ad := range add.Array() {
			fmt.Println(ad)
			addedResult = append(addedResult, ad.String())
		}
	}
	fmt.Println(addedResult)

	//从前端获得数据，并解包放入resp中

	//	data := this.Ctx.Input.RequestBody
	//	fmt.Println(&resp)
	//	fmt.Println(string(resp["ref"]))
	//fmt.Println()

	//fmt.Println("start kafka send")
	//syncProducer(Address,json)

	this.Ctx.WriteString("chenggong")
	fmt.Println("over ---")
}

/*
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
    topic := _topic
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
}*/
