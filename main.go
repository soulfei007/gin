package main

import (
	"fmt"

	"gin-learn/pkg/setting" //导入setting 包，会自动执行init()函数，给变量赋值
	"net/http"

	"gin-learn/models"
	"gin-learn/pkg/gredis"
	"gin-learn/pkg/logging"
	"gin-learn/routers"
)

func main() {
	setting.Steup()
	models.Setup()

	logging.Setup()
	gredis.Setup()

	router := routers.InitRouter()

	t := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	t.ListenAndServe()
}
