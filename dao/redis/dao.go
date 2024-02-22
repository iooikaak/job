package redis

import (
	"github.com/iooikaak/frame/cache/redis/v8"
	"github.com/iooikaak/job/config"
)

type Dao struct {
	redis *redis.Client
}

//New .
func New(cfg *config.Config) (d *Dao) {
	d = &Dao{
		redis: redis.New(cfg.Redis),
	}
	return
}

// Stop close the resource.
func (d *Dao) Stop() error {
	return d.redis.Close()
}
