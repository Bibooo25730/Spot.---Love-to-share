package admin_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/response"
	"index_Demo/utils/userUtil"
	"net/http"
)

func ReviewFeedback(ctx *gin.Context) {
	if !userUtil.IsAdmin(ctx) {
		ctx.JSON(http.StatusUnauthorized, response.New("Unauthorized", nil))
		return
	}
	QueryFdBack, pagination, err := userUtil.QueryFdBack(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
		"list":     QueryFdBack,
		"total":    pagination.Total,
		"pages":    pagination.Total / int64(pagination.PageSize),
		"pageNum":  pagination.PageNum,
		"pageSize": pagination.PageSize,
	}))

}
