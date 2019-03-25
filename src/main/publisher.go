package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main(){
	opts := mqtt.NewClientOptions()
	opts.AddBroker("test.mosquitto.org:1883")
	
	publisher := mqtt.NewClient(opts)
	token := publisher.Connect()
	
	for !token.WaitTimeout(3 * time.Second){
		
	}
	
	if token.Error() != nil {
		fmt.Println("token err:",token.Error())
	}
	fmt.Println("publisher is here...")
	
	for{
	fmt.Println("Please input what you want to send,input exit to exit..")
	msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
	
	if err!=nil{
		fmt.Println("err in read..",err)
	}
	
	msg = strings.Trim(msg, "\r\n")
	if msg=="exit"{
		fmt.Println("You are exited..")
		break
	}
	
	fmt.Println("you input:",msg)
	
	token = publisher.Publish("linyujia/testtopic", 0, false, msg)
	if token.Error() != nil{
		fmt.Println("publish error:",token.Error())
	}
	fmt.Println("publish done...")
	
	
	}
	
	
}