package redis

import (
	"testing"
	"time"
)

func TestDao_LockSupplier(t *testing.T) {
	a := make(chan struct{})
	r, err := d.LockShow(ctx, 114)
	t.Logf("%v-----%v", r, err)
	ti := time.NewTicker(10 * time.Second)
	for i := range ti.C {
		url, _ := config.GetPxqUrl()
		t.Logf("%#v-----%v---%v", i, config.M.Get("pxq_url"), url)
	}
	<-a
}

func TestDao_GetSupplierLock(t *testing.T) {
	r, err := d.GetShowLock(ctx, 111)
	defer func() {
		if r {
			t.Logf("fffff")
			_ = d.ReleaseShowLock(ctx, 111)
		}
	}()
	t.Logf("%v-----%v", r, err)
	if r {
		t.Logf("dddddd")
		_ = d.ReleaseShowLock(ctx, 111)
		r = false
	}
	t.Logf("%v-----%v", r, err)
}

func TestDao_ReleaseSupplierLock(t *testing.T) {
	err := d.ReleaseShowLock(ctx, 111)
	t.Logf("%v", err)
}
