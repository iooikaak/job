package net_edu

import (
	"github.com/iooikaak/frame/gorm"
	"github.com/iooikaak/job/config"
)

const MaxPrice = 999999

type Dao struct {
	Db *gorm.Engine
}

func New(cfg *config.Config) (d *Dao) {
	d = &Dao{}
	d.Db = gorm.New(cfg.DB)
	return
}

func (d *Dao) Stop() (err error) {
	return d.Db.Close()
}
