package main
import (
    "github.com/Shopify/sarama"
    "time"
    "log"
    "fmt"
    "os"
   // "os/signal"
    //"sync"
)
 
var Address = []string{"47.97.248.41:9092"}
 
func main()  {
    syncProducer(Address)
    //asyncProducer1(Address)
}
 
//同步消息模式
func syncProducer(address []string)  {
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
    srcValue := "sync: this is a zuolonglong. index=%d"
    for i:=0; i<10; i++ {
        value := fmt.Sprintf(srcValue, i)
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
    }
}