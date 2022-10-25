package initialize

import (
	"encoding/json"
	"go-api-base/config"
	"os"
)

func Config() {
	var cfg = struct {
		DB     *config.DBConfig
		Server *config.ServerConfig
	}{}
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic("配置文件未找到:" + err.Error())
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic("配置文件解析失败" + err.Error())
	}
	config.Server = cfg.Server
	config.DB = cfg.DB
}
