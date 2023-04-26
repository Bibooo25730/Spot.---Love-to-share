package serverUtils

import "github.com/gin-gonic/gin"

func GetRealIP(ctx *gin.Context) string {
	return ctx.ClientIP()
	//ctx.RemoteIP()
}
