package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"net/http"
	"time"
)

func FeedBack(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)
	feedBackSql := dal.Feedback
	fdBackJson := model.Feedback{}
	err := ctx.BindJSON(&fdBackJson)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("系统内部出错 请联系管理员", nil))
		return
	}
	seed := model.Feedback{
		FUser: user.Username,
		FMsg:  fdBackJson.FMsg,
		FTime: time.Now(),
	}
	err = feedBackSql.WithContext(ctx).Create(&seed)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("反馈失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.New("反馈成功", seed))
}
