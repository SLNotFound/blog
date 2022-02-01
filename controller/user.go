package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1 参数校验
	p := new(models.SignUpParam)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": RemoveTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 2 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func LoginHandler(c *gin.Context) {
	// 1 参数校验
	p := new(models.LoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": RemoveTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 2 业务处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录失败",
		})
		return
	}
	// 3 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
