package grammar_test

import (
	"fmt"
	"github.com/monnand/goredis"
	"log"
	"testing"
	"time"
)

type Message struct {
}

func Test_String(t *testing.T) {
	//redis基本操作
	var client goredis.Client
	//设置基本redis信息
	client.Addr = "127.0.0.1:6379"
	client.Db = 1
	//string操作
	client.Set("test", []byte("hello"))
	res, _ := client.Get("test")
	fmt.Println(string(res))
	client.Del("test")
}

func Test_List(t *testing.T) {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	client.Db = 1

	list := []string{"a", "b", "c", "d"}
	for _, v := range list {
		client.Rpush("list", []byte(v))
	}
	res, err := client.Lrange("list", 0, 3)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range res {
		fmt.Println(i, ":", string(v))
	}

	client.Del("list")
}

func Test_Pub(t *testing.T) {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	client.Db = 1

	sub := make(chan string, 1)
	sub <- "foo"
	messages := make(chan goredis.Message, 0)
	go client.Subscribe(sub, nil, nil, nil, messages)

	time.Sleep(10 * 1000 * 1000)
	client.Publish("foo", []byte(("bar")))

	msg := <-messages
	println("received from", msg.Channel, " message:", string(msg.Message))

	close(sub)
	close(messages)
}
