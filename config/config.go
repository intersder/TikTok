package config

import (
	"encoding/json"
	"os"
)

var Config = newConfig()

func newConfig() *config {

	bytes, err := os.ReadFile("config/config_ch.json")
	if err != nil {
		panic("config file read fail. " + err.Error())
		return nil
	}
	config := new(config)

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		panic("config file unmarshal fail. " + err.Error())
		return nil
	}
	return config
}

type config struct {
	DBUrl        string // 数据库连接字符串
	ShowSql      bool   // 是否显示sql
	LogFile      string // 日志文件
	RDBUrl       string // redis连接字符串
	RDBPassword  string // redis密码
	CosRegion    string // cos区域
	CosBucket    string // cos bucket名称
	CosSecretId  string // cos secretId
	CosSecretKey string // cos secretKey
}
