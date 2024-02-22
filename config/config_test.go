package config

import (
	"encoding/json"
	"testing"
)

func TestInit(t *testing.T) {
	if err := ApolloInit(); err != nil {
		panic(err)
	}
	bs, _ := json.Marshal(Conf)
	t.Log(string(bs))
}
