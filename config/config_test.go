package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {

	LoadConfig("../conf/config.json")

	if GetIsInit() != true {
		t.Error(" 初始化失败")
	}
	if GetServer().Port == 8181 {
		t.Log("test port faild")
	}

}
