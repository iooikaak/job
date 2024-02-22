package channel_center

//
//import (
//	"context"
//	"testing"
//
//	"github.com/iooikaak/frame/snowflake"
//
//	model "git.juqitech.com/go/drama/common/model/mysql/channel_center"
//)
//
//func TestGetAuthUser(t *testing.T) {
//	a, b := d.GetAuthUser(context.Background(), 374474119980277479)
//	t.Logf("---%v---%v---", a, b)
//}
//
//func TestGetAuthUserN(t *testing.T) {
//	db := d.Db.Context(context.Background())
//	a, b := d.GetAuthUserN(db.Debug(), 374474119980277479)
//	t.Logf("---%v---%v---", a, b)
//}
//
//func TestInsertAuthUser(t *testing.T) {
//	db := d.Db.Context(context.Background())
//	id, _ := snowflake.NewID()
//	a := d.InsertAuthUser(db.Debug(), &model.AuthUser{
//		ID:      int64(id),
//		Account: "test_user",
//		Name:    "test_name",
//		Pwd:     "test_pwd",
//	})
//	t.Logf("---%v---", a)
//}
//
//func TestUpdateAuthUser(t *testing.T) {
//	db := d.Db.Context(context.Background())
//	//id, _ := snowflake.NewID()
//	var m = &model.AuthUser{
//		//ID:      111,
//		Name:    "fasdfds",
//		Account: "sdfdfd",
//	}
//	a := d.UpdateAuthUser(db.Debug(), m, &model.AuthUser{
//		//ID:   1111,
//		Name: "fasdfds",
//	})
//	t.Logf("---%v---", a)
//}
//
//func TestGetAuthUsers(t *testing.T) {
//	a, b := d.GetAuthUsers(context.Background(), "test_name")
//	t.Logf("---%#v---%v---", a, b)
//}
