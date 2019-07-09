package dashboard

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func Get(c *gin.Context) {
	data := map[string]interface{}{
		"k": "测试",
	}

	c.JSONP(http.StatusOK, data)
}
