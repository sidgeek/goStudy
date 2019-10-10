package ginTest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("", func(c *gin.Context) {

		c.String(http.StatusOK, "Hello GoLang!")
	})
	router.Run(":8083")
}