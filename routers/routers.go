package routers

import (
	"blog/controller"
	"blog/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controller.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})
	return r
}
