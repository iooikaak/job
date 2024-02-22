package service

//
//import (
//	"context"
//	"os"
//
//	"git.juqitech.com/go/drama/common/kdniao"
//
//	"testing"
//
//	"github.com/iooikaak/job/conf"
//)
//
//var srv *Service
//
//func TestMain(t *testing.M) {
//	//_ = flag.Set("conf", "../build/config.yaml")
//	if err := conf.ApolloInit(); err != nil {
//		panic(err)
//	}
//	srv = New(conf.Conf)
//	kdniao.Init(conf.Conf.ExpressTracing)
//	os.Exit(t.Run())
//}
//
//func TestExpressTracing(t *testing.T) {
//	a, b := kdniao.GetTracing(context.Background(), "ZTO", "75833211124597", "")
//	t.Log(a, b)
//}
