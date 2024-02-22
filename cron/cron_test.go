package cron

import (
	"flag"
	"github.com/iooikaak/job/config"
	"os"

	"testing"
)

func TestMain(t *testing.M) {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}

	os.Exit(t.Run())
}
