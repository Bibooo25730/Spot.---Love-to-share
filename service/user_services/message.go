package user_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/app/request"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/middleware/auth"
	"net/http"
	"time"
)

func Message(ctx *gin.Context) {
	user := auth.CurrentUser(ctx)

	msgRes := request.MessageRes{}
	err := ctx.ShouldBindJSON(&msgRes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("绑定数据失败", err))
		return
	}

	contentReq := model.Message{
		Uname:    &user.Username,
		CreateAt: time.Now(),
		Content:  &msgRes.Content,
	}
	content := dal.Message
	err = content.WithContext(ctx).Create(&contentReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("留言失败", err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.New("留言成功", msgRes.Content))
}
