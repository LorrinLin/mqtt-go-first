package main

import (
	"time"
	"bufio"
	"fmt"
	"os"
	"sync"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var(
	wg sync.WaitGroup
)

func main(){
	uri := "iot.eclipse.org:1883"
	
	topic := ""
	if topic == "" {
		topic = "test"
	}

	wg.Add(1)
	listen(uri,topic)
	
	client := connect("pub",uri)
	
	fmt.Println("Please input what you want to send..")
	msg,err := bufio.NewReader(os.Stdin).ReadString('\n')
	
//	timer := time.NewTicker(1 * time.Second)
//	for t := range timer.C {
//		client.Publish(topic, 0, false, "Eric: "+t.String())
//	}

	if err!=nil{
		fmt.Println("err in read..",err)
	}
	fmt.Println("you input:",msg)
	client.Publish(topic, 0, false, msg)
	
	wg.Wait()
}

func listen(uri string, topic string){
	
	fmt.Println("in listen...")
	client := connect("sub",uri)
	client.Subscribe(topic, 0, func (client mqtt.Client, msg mqtt.Message){
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		wg.Done()
	})
	
}

func connect(clientId string, uri string) mqtt.Client{
	fmt.Println("in connect...")
	opt := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opt)
	token := client.Connect()
	
	for !token.WaitTimeout(3 * time.Second){
		
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
