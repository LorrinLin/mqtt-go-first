package main

import (
	"fmt"
	"os"
	"bufio"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main(){
	opts := mqtt.NewClientOptions()
	opts.AddBroker("iot.eclipse.org:1883")
	
	publisher := mqtt.NewClient(opts)
	token := publisher.Connect()
	
	if token.Error() != nil {
		fmt.Println("token err:",token.Error())
	}
	fmt.Println("publisher is here...")
	fmt.Println("Please input what you want to send..")
	msg,err := bufio.NewReader(os.Stdin).ReadString('\n')
	
	if err!=nil{
		fmt.Println("err in read..",err)
	}
	fmt.Println("you input:",msg)
	
	token = publisher.Publish("my/test", 0, false, msg)
	if token.Error() != nil{
		fmt.Println("publish error:",token.Error())
	}
	fmt.Println("publish done...")
	
}