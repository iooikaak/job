package main

import (
	"context"
	"flag"
	"github.com/iooikaak/frame/bootstrap"
	"github.com/iooikaak/frame/common/kdniao"
	"github.com/iooikaak/job/config"
	"github.com/iooikaak/job/consumer"
	"github.com/iooikaak/job/service"

	//"github.com/iooikaak/frame/drama/common/handler/client"
	"github.com/iooikaak/frame/sync/errgroup"
	"github.com/iooikaak/job/cron"
)

var srv *service.Service

func main() {
	flag.Parse()
	//远程读取apollo配置文件
	//初始化apollo
	if err := config.ApolloInit(); err != nil {
		panic(err)
	}
	// 启动框架
	app := bootstrap.New(bootstrap.Config{
		Log:         config.Conf.Log,
		Tracer:      config.Conf.Tracer,
		GinServer:   config.Conf.GinServer,
		MicroServer: config.Conf.MicroServer,
		ServiceName: config.Conf.ServiceName,
	})
	srv = service.New(config.Conf)
	app.Init(
		bootstrap.BeforeStart(func() {
			kdniao.Init(config.Conf.ExpressTracing)
		}),
		bootstrap.AfterStart(func() {
			eg := new(errgroup.Group)
			//消费消息队列
			eg.Go(func(ctx context.Context) error {
				consumer.Init(srv)
				return nil
			})
			//定时任务
			eg.Go(func(ctx context.Context) error {
				cron.Init(srv)
				return nil
			})
			_ = eg.Wait()
		}),
		bootstrap.AfterStop(func() {
			eg := new(errgroup.Group)
			eg.Go(func(ctx context.Context) error {
				//消息队列平稳停止
				consumer.Stop()
				return nil
			})
			_ = eg.Wait()
		}),
		//bootstrap.MicroService(func(s *micro.Service) {
		//	client.InitRpc(s.Client())
		//}),
	)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
