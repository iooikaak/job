package channel_center

import (
	"github.com/iooikaak/job/config"
	"os"
	"testing"
)

var d *Dao

func TestMain(m *testing.M) {
	//_ = flag.Set("conf", "../../../build/config.yaml")
	if err := config.ApolloInit(); err != nil {
		panic(err)
	}
	d = New(config.Conf)
	os.Exit(m.Run())
}
