package cron

import (
	"github.com/iooikaak/job/service"
	"github.com/robfig/cron"
)

var srv *service.Service

func Init(s *service.Service) {
	srv = s
	c := cron.New()
	_ = c.AddFunc("* * * * * *", show)
	// crawl
	//show()
	//_ = c.AddFunc("@every 5s", show)
	c.Start()
}
