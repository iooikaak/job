package net_edu

import (
	"context"
	"testing"

	"github.com/iooikaak/frame/snowflake"

	model "github.com/iooikaak/job/common/model/mysql/net_edu"
)

func TestGetNetEdu(t *testing.T) {
	a, b := d.GetNetEdu(context.Background(), 435585656115185520)
	t.Logf("---%v---%v---", a, b)
}

func TestGetNetEduN(t *testing.T) {
	db := d.Db.Context(context.Background())
	a, b := d.GetNetEduN(db.Debug(), 374474119980277479)
	t.Logf("---%v---%v---", a, b)
}

func TestInsertNetEdu(t *testing.T) {
	db := d.Db.Context(context.Background())
	id, _ := snowflake.NewID()
	a := d.InsertNetEdu(db.Debug(), &model.NetEdu{
		ID:   int64(id),
		Body: "0123",
	})
	t.Logf("---%v---", a)
}

func TestUpdateNetEdu(t *testing.T) {
	db := d.Db.Context(context.Background())
	//id, _ := snowflake.NewID()
	var m = &model.NetEdu{
		//ID:      111,
		Body: "0123",
	}
	a := d.UpdateNetEdu(db.Debug(), m, &model.NetEdu{
		//ID:   1111,
		Body: "0123",
	})
	t.Logf("---%v---", a)
}

func TestGetNetEdus(t *testing.T) {
	a, b := d.GetNetEdus(context.Background(), "test_name")
	t.Logf("---%#v---%v---", a, b)
}
