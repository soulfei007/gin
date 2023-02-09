package main

import (
	"fmt"

	"gin-learn/pkg/setting" //导入setting 包，会自动执行init()函数，给变量赋值
	"net/http"

	"gin-learn/routers"
)

func main() {
	router := routers.InitRouter()

	t := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	t.ListenAndServe()
}
