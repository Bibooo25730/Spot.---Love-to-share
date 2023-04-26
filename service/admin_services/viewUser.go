package admin_services

import (
	"github.com/gin-gonic/gin"
	"index_Demo/gen/response"
	"index_Demo/utils/userUtil"
	"net/http"
)

func ViewUserList(ctx *gin.Context) {
	//该用户是否为管理员
	if !userUtil.IsAdmin(ctx) {
		ctx.JSON(http.StatusUnauthorized, response.New("Unauthorized", nil))
		return
	}
	queryUsers, pagination, err := userUtil.QueryUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, response.New("查询成功", gin.H{
		"list":     queryUsers,
		"total":    pagination.Total,
		"pages":    pagination.Total / int64(pagination.PageSize),
		"pageNum":  pagination.PageNum,
		"pageSize": pagination.PageSize,
	}))
}
