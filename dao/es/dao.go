package es

import (
	"github.com/iooikaak/job/config"
	"time"

	"github.com/iooikaak/frame/elastic"
	"github.com/iooikaak/frame/xlog"

	"github.com/iooikaak/frame/elasticsearch"
)

const (
	MaxEsRetries    = 5
	BackOffDelayMin = 100 * time.Millisecond
	BackOffDelayMax = 1 * time.Second
)

type Dao struct {
	es   *elastic.Client
	conf *elasticsearch.ElasticConfig
}

func New(cfg *config.Config) (d *Dao) {
	instance, err := elasticsearch.New(cfg.Elastic,
		elastic.SetErrorLog(xlog.Logger()), elastic.SetRetrier(elastic.NewBackoffRetrier(&WrapperRetries{
			backOffDelayMax: BackOffDelayMax,
			backOffDelayMin: BackOffDelayMin,
			retryNum:        MaxEsRetries,
		})),
	)
	if err != nil {
		panic(err)
	}
	return &Dao{es: instance, conf: cfg.Elastic}
}

type WrapperRetries struct {
	retryNum                         int
	backOffDelayMin, backOffDelayMax time.Duration
}

func (w WrapperRetries) Next(retry int) (time.Duration, bool) {
	if retry > w.retryNum {
		return 0, false
	}
	return backOff(retry, w.backOffDelayMin, w.backOffDelayMax), true
}

func backOff(attempt int, min time.Duration, max time.Duration) time.Duration {
	d := time.Duration(attempt*attempt) * min
	if d > max {
		d = max
	}
	return d
}

func (d *Dao) Stop() {
	d.es.Stop()
}
