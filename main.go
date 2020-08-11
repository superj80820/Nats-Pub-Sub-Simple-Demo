package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

func blockProcess() {
	select {}
}

func main() {
	// 連接至Nats
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		fmt.Println(err)
	}

	// 使用Channel對Nats來subscribe 'foo'這個topic
	ch := make(chan *nats.Msg, 64)
	_, err = nc.ChanSubscribe("foo", ch)

	// Publish 'Hello World'至'foo' topic
	nc.Publish("foo", []byte("Hello World"))

	// 透過Channel來接收Publish端傳送的訊息
	msg := <-ch
	fmt.Printf("Received a message: %s\n", string(msg.Data))

	blockProcess()
}
