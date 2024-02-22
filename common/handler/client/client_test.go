package client

//import (
//	"context"
//	"os"
//	"testing"
//
//	"git.juqitech.com/go/demo/api/conf"
//	"github.com/iooikaak/frame/micro"
//
//	proto "git.juqitech.com/go/proto/demoAdmin"
//)
//
//func TestMain(m *testing.M) {
//	c := micro.New(conf.Conf.MicroServer)
//	_ = c.Initialize()
//	InitRpc(c.Client())
//	os.Exit(m.Run())
//}
//
//func TestDemoAdminOperate(t *testing.T) {
//	ctx := context.Background()
//	input := &proto.OperateInput{}
//	a, b := DemoAdmin().Operate(ctx, input)
//	t.Logf("----%v----%v----", a, b)
//}
