package routers

import (
	"blog/controller"
	"blog/logger"
	"blog/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/hello", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录用户，判断请求头中是否有 有效的JWT
		c.String(http.StatusOK, "hello gin")
	})
	return r
}
