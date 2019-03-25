package main

import (
	"time"
	"bufio"
	"fmt"
	"os"
	"sync"
	"strings"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var(
	wg sync.WaitGroup
)

func main(){
	uri := "test.mosquitto.org:1883"
	
	topic := ""
	if topic == "" {
		topic = "linyujia/test"
	}

	fmt.Println("Please input what you want to send,input 'exit' to exit..")
	
	publisher := connect("pub", uri)
	go listen(uri,topic)
	wg.Add(1)
		
	for{
		msg,err := bufio.NewReader(os.Stdin).ReadString('\n')
		msg = strings.Trim(msg, "\r\n")
		if msg == "exit"{
			fmt.Println("you are exited..")
			break
		}

		if err!=nil{
			fmt.Println("err in read..",err)
		
		}
		
		token := publisher.Publish(topic, 0, false, msg)
		if token.Error() != nil{
			fmt.Println("err in publish - ",token.Error())
			break
		}

	}
	fmt.Println("----------byebye---------")
}

func listen(uri string, topic string){
	
	fmt.Println("in listen...")
	consumer := connect("sub",uri)
	consumer.Subscribe(topic, 0, func (client mqtt.Client, msg mqtt.Message){
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))

	})
	
}

func connect(clientId string, uri string) mqtt.Client{
	fmt.Println("in connect...")
	opt := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opt)
	token := client.Connect()
	
	for !token.WaitTimeout(1 * time.Second){
		
	}
	if err := token.Error(); err != nil{
		fmt.Println("err in connect--token,",err)
	}
	
	return client
	
}

func createClientOptions(clientId string, uri string) *mqtt.ClientOptions{
	fmt.Println("in createClientOptions...")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetClientID(clientId)
	return opts
}
