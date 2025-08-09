package Subscriber

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/FACELESS-GOD/LoggerInjest.git/Package/Processor"
	"github.com/redis/go-redis/v9"
)

type SubInt interface {
	Subscribe()
}

type SubStruct struct {
	Rdb        *redis.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
	RDBChannel string
	RDBPubSub  *redis.PubSub
	Proc       Processor.ProcStruct
}

func Sub() SubStruct {

	sub := SubStruct{}

	channelName := os.Getenv("Logger_Channel")
	sub.RDBChannel = channelName

	redisOps := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	rdb := redis.NewClient(&redisOps)
	sub.Rdb = rdb

	ctx, cle := context.WithTimeout(context.Background(), time.Second*10)
	sub.Ctx = ctx
	sub.Cancel = cle

	pubSub := sub.Rdb.Subscribe(sub.Ctx, sub.RDBChannel)
	sub.RDBPubSub = pubSub

	proc := Processor.Proc()
	sub.Proc = proc

	return sub
}

func (Sub *SubStruct) Subscribe() {

	for {
		msg, err := Sub.RDBPubSub.ReceiveMessage(Sub.Ctx)

		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, msg.Payload)

		Sub.Proc.Process(msg.Channel, msg.Payload)
	}

	defer Sub.RDBPubSub.Close()
}
