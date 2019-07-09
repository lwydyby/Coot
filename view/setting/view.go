package setting

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
