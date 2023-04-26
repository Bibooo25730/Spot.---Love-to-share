package user_services

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"index_Demo/dao/redisServer"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/response"
	"index_Demo/utils/email"
	"net/http"
	"time"
)

type RegEmail struct {
	Email string `json:"email" binding:"required"`
}

func RegEmailCode(ctx *gin.Context) {
	regEmail := RegEmail{}
	err := ctx.BindJSON(&regEmail)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("系统出错联系管理员", err))
		return
	}
	u := dal.User
	user, _ := u.WithContext(ctx).Where(u.Uemail.Eq(regEmail.Email)).First()
	if user != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("邮箱已存在", nil))
		return
	}
	code, emailErr := email.SendMail(regEmail.Email)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("验证码发送失败 联系系统管理员", emailErr))
		return
	}
	redisServer.Set(regEmail.Email, *code, time.Duration(viper.GetInt("redis.emailTime"))*time.Minute)

	ctx.JSON(http.StatusOK, response.New("验证码发送成功", code))
}
