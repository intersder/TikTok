package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
	"vibrato/model"
	"vibrato/sqls"
)

func init() {
	// gorm配置
	gormConf := &gorm.Config{}

	// 初始化日志
	if file, err := os.OpenFile(Config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
		if Config.ShowSql {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			})
		}
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
	}

	// 连接数据库
	if err := sqls.Open(Config.DBUrl, gormConf, 32, 128, model.Models...); err != nil {
		logrus.Error(err)
	}
	// 连接redis
	sqls.OpenRedisClient(Config.RDBUrl, Config.RDBPassword, 0)

}

func main() {

	//user := model.User{
	//	Name: "dousheng",
	//}
	//token, err := services.GenerateToken(&user)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(token)
	//user, err := services.GetUserByToken("dabc3eadfa9b4564ab92b80a9cd989e4")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//bytes, _ := json.Marshal(user)
	//
	//log.Println(string(bytes))
	//u := uuid.NewV4()
	//
	//print(strings.ReplaceAll(u.String(), "-", ""))

	//result, err := sqls.RDB().Set(context.Background(), "test", "test666", time.Second*10).Result()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(result)
	//
	//res, err := sqls.RDB().Get(context.Background(), "test").Result()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//log.Println(res)

	//user, err := services.UserService.Register("admin3", "admin3", "admin3")
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}
	//bytes, err := json.Marshal(user)
	//if err != nil {
	//	return
	//}
	//fmt.Printf(string(bytes))

	r := gin.Default()
	initRouter(r)
	err := r.Run()
	if err != nil {
		logrus.Error("服务启动失败", err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
