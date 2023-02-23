package app

import (
	"github.com/gin-gonic/gin"

	"gin-learn/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (c Gin) Response(httpCode, errCode int, data interface{}) {
	c.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}
