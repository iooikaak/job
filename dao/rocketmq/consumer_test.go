package rocketmq

import (
	"fmt"
	"testing"

	"github.com/iooikaak/frame/mq/rocketmq"
)

func TestConsumer(t *testing.T) {
	a := make(chan struct{})
	selector := rocketmq.ConsumerMessageSelector{}
	err := d.Consumer.RegisterHandle(ConsumerTopicTest, selector, CallBack)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = d.Consumer.Start()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	<-a
}

func CallBack(msgs []*rocketmq.MessageExt, ctx *rocketmq.ConsumeConcurrentlyContext) (rocketmq.ConsumeResult, error) {
	for _, msg := range msgs {
		fmt.Printf("Consumer msg: %#v body: %+v queue1111: %d", msg, string(msg.Body), msg.Queue.QueueId)
	}
	return rocketmq.ConsumeSuccess, nil
}
