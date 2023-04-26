package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"index_Demo/dao/redisServer"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求头中的 token
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.New("Unauthorized", nil))
			return
		}
		// 根据 token 获取用户数据
		user, err := getUserByToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.New("Unauthorized", nil))
			return
		}

		// 将用户数据保存到上下文中
		ctx.Set("user", user)
		// 继续处理请求
		ctx.Next()
	}
}

func getUserByToken(token string) (*model.User, error) {
	//查看rides中是否存在token 如果存在，返回用户信息 如果不存在，返回错误
	jsonStr, err := redisServer.Get(fmt.Sprintf("user_token_%s", token))
	if err != nil {
		return nil, err
	}
	u := model.User{}
	err = json.Unmarshal([]byte(jsonStr), &u)
	return &u, err
}
