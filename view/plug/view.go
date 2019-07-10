package plug

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "plug.html", gin.H{
		"title": "Main website",
	})
}
