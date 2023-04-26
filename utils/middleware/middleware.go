package middleware

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/response"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, response.New("系统内部错误，请联系管理员", nil))
				log.Printf("Panic: %v\n", err)
				bytes := make([]byte, 1024*24)
				runtime.Stack(bytes, false)
				log.Println(string(bytes))
			}
		}()
		c.Next()
	}
}

func DeviceType() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.Request.UserAgent()
		if strings.Contains(userAgent, "Mobile") {
			c.Set("device_type", "mobile")
		} else {
			c.Set("device_type", "desktop")
		}
		c.Next()
	}
}

func HandleDeviceType(c *gin.Context) {
	deviceType, exists := c.Get("device_type")
	if !exists {
		c.String(500, "Failed to detect device type")
	} else if deviceType == "mobile" {
		c.String(200, "This is a mobile device")
	} else {
		c.String(200, "This is a desktop device")
	}
}
