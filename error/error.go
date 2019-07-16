package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Check(e error, tips string) {
	if e != nil {
		panic(e)
		fmt.Println(tips)
	}
}

func ErrSuccess(data []map[string]interface{}) map[string]interface{} {
	return gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	}
}
func ErrSuccessCustom(code int, msg string, data []map[string]interface{}) map[string]interface{} {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
func ErrSuccessNull() map[string]interface{} {
	return gin.H{
		"code": 200,
		"msg":  "success",
	}
}

func ErrFailFileType() map[string]interface{} {
	return gin.H{
		"code": 101,
		"msg":  "异常文件格式",
	}
}
