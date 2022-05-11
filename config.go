package main

var Config = newConfig()

func newConfig() *config {
	return &config{
		DBUrl:       "root:root@tcp(192.168.146.130:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local",
		ShowSql:     true,
		LogFile:     "D:/logs/vibrato.log",
		RDBUrl:      "192.168.146.130:6379",
		RDBPassword: "",
	}
}

type config struct {
	DBUrl       string // 数据库连接字符串
	ShowSql     bool   // 是否显示sql
	LogFile     string // 日志文件
	RDBUrl      string // redis连接字符串
	RDBPassword string // redis密码
}
