package main

import(
	"sync"
	"fmt"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	wg sync.WaitGroup
)

func main(){
	opts := mqtt.NewClientOptions()
	opts.AddBroker("test.mosquitto.org:1883")
	
	
	go process(opts)
	
	wg.Add(1)
	wg.Wait()
	
	
}
func process(opts *mqtt.ClientOptions){	
	consumer := mqtt.NewClient(opts)
	token := consumer.Connect()
	for !token.WaitTimeout(3 * time.Second){
		
	}
	
	if token.Error() != nil {
		fmt.Println("token err:",token.Error())
	}
	
	consumer.Subscribe("linyujia/testtopic", 0, func(client mqtt.Client, msg mqtt.Message){
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		
	})
	fmt.Println("consumer is here..")
	fmt.Println("waiting for your message..")
	
	}
	