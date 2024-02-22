package net_edu

import (
	"github.com/iooikaak/job/config"
	"os"
	"testing"
)

var d *Dao

func TestMain(m *testing.M) {
	//_ = flag.Set("conf", "../../../config/job.json")
	//if err := config.ApolloInit(); err != nil {
	//	panic(err)
	//}
	config.Init()
	d = New(config.Conf)
	os.Exit(m.Run())
}
