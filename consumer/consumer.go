package consumer

import (
	"context"
	"time"

	"github.com/iooikaak/frame/sync/errgroup"
	"github.com/iooikaak/job/service"
)

var srv *service.Service

func Init(s *service.Service) {
	srv = s
	eg := new(errgroup.Group)
	eg.Go(func(ctx context.Context) error {
		srv.StartShow()
		return nil
	})
	_ = eg.Wait()
}

func Stop() {
	// 先批量关闭消费者，然后关闭资源
	eg := new(errgroup.Group)
	eg.Go(func(ctx context.Context) error {
		srv.StopShow()
		return nil
	})
	_ = eg.Wait()
	time.Sleep(20 * time.Second)
	srv.Stop()
}
