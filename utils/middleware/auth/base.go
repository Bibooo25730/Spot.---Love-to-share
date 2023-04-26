package auth

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/orm/model"
)

func CurrentUser(ctx *gin.Context) *model.User {
	value, exists := ctx.Get("user")
	if exists {
		return value.(*model.User)
	}
	return nil
}
