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
