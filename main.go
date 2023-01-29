package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main(){
	
	url := "nats://0.0.0.0:50560"
// Connect to a server
	nc, err := nats.Connect(url, nats.UserInfo("local", "roSpw9c9O4fwKNBpZ1tOPiNtilmQ9KF6"))
	if err != nil {
		logrus.Fatal(err)
	}
	defer nc.Close()
	sub, _ := nc.Subscribe("hello.john", func(msg *nats.Msg){
		fmt.Println("Subject : ? , Data : ?\n", msg.Subject, string(msg.Data))
	})

	time.Sleep(100 * time.Second)
	sub.Unsubscribe()
}