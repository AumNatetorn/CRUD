package configs_test

import (
	"CRUD/configs"
	"encoding/json"
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetConfig(t *testing.T) {
	configs.Init(".")

	conf := configs.GetConfig()

	assert.Equal(t, "CRUD", conf.App.Name)

	b, _ := json.Marshal(conf)
	fmt.Println(string(b))
}
