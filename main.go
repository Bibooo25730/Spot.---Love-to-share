package main

import (
	"fmt"
	logs "github.com/danbai225/go-logs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"index_Demo/config"
	mysql "index_Demo/dao/mysql"
	"index_Demo/dao/redisServer"
	"index_Demo/router"
)

func main() {
	//加载config文件
	err := config.InfoConfig()
	if err != nil {
		logs.Err(err.Error())
		return
	}
	//err = config.LoadConfig()
	//if err != nil {
	//	logs.Err(err)
	//	return
	//}

	//连接sql
	err = mysql.Init()
	if err != nil {
		logs.Err(err.Error())
		return
	}
	//连接redis
	redisServer.Init()

	//初始化Gin
	g := gin.Default()
	router.Router(g)
	//http://patorjk.com/software/taag/#p=display&h=0&f=Ogre&t=stephen
	//logs.Info("\n      _                 _                  \n ___ | |_   ___  _ __  | |__    ___  _ __  \n/ __|| __| / _ \\| '_ \\ | '_ \\  / _ \\| '_ \\ \n\\__ \\| |_ |  __/| |_) || | | ||  __/| | | |\n|___/ \\__| \\___|| .__/ |_| |_| \\___||_| |_|\n                |_|                        \n")
	logs.Info("\n     _                         \n  __| |  ___  _ __ ___    ___  \n / _` | / _ \\| '_ ` _ \\  / _ \\ \n| (_| ||  __/| | | | | || (_) |\n \\__,_| \\___||_| |_| |_| \\___/ \n                               \n")
	err = g.Run(fmt.Sprintf(":%d", viper.GetUint16("server.port")))
	if err != nil {
		logs.Err(err.Error())
	}

}
