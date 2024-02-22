package service

import (
	"github.com/iooikaak/frame/stat/prom"
	"github.com/iooikaak/frame/xlog"
	"github.com/iooikaak/job/config"
	"github.com/iooikaak/job/dao/es"
	"github.com/iooikaak/job/dao/mysql/channel_center"
	"github.com/iooikaak/job/dao/redis"
	"github.com/iooikaak/job/dao/rocketmq"
)

type Service struct {
	mq    *rocketmq.Dao
	redis *redis.Dao
	es    *es.Dao
	db    *channel_center.Dao
	prom  *prom.Prom
}

func New(cfg *config.Config) *Service {
	srv := &Service{
		mq:    rocketmq.New(cfg),
		redis: redis.New(cfg),
		es:    es.New(cfg),
		db:    channel_center.New(cfg),
		prom:  prom.New().WithCounter("orderAdapterCount", []string{"operation", "topic"}),
	}
	return srv
}

func (s *Service) Stop() {
	err := s.db.Stop()
	if err != nil {
		xlog.Errorf("Stop db failed err: %v", err)
	}
	s.es.Stop()
	err = s.redis.Stop()
	if err != nil {
		xlog.Errorf("Stop redis failed err: %v", err)
	}
	err = s.mq.Producer.Stop()
	if err != nil {
		xlog.Errorf("Stop producer failed err: %v", err)
	}
}
