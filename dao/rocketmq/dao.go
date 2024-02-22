package rocketmq

import (
	"github.com/iooikaak/frame/mq/rocketmq"
	"github.com/iooikaak/job/config"
)

type Dao struct {
	Producer *rocketmq.Producer
	Consumer *rocketmq.PushConsumer
}

//New .
func New(cfg *config.Config) (d *Dao) {
	d = &Dao{}
	if p, err := rocketmq.NewProducer(cfg.RocketMq); err == nil {
		d.Producer = p
	} else {
		panic(err)
	}
	if c, err := rocketmq.NewPushConsumer(cfg.RocketMq, ConsumerGroupTest); err == nil {
		d.Consumer = c
	} else {
		panic(err)
	}
	return
}
