package config

import "github.com/spf13/viper"

// 服务名
var ServerName string

// 模式, local, dev, pro
var Mode string

// 主服务端口
var HttpPort string

var Cfg *viper.Viper

func InitConfig() {
	Cfg = viper.New()
	var filePath string
	switch Mode {
	case "local":
		filePath = "G:\\go_workspace\\GOPATH\\src\\go-learn\\projects\\online-edu\\config\\pro.yaml"
	case "pro":
		filePath = "/data/projects/online-edu/config/pro.yaml"
	}
	Cfg.SetConfigFile(filePath)
	if e := Cfg.ReadInConfig(); e != nil {
		panic(e)
	}
}
