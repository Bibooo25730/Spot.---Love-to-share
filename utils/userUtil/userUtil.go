package userUtil

import (
	"github.com/gin-gonic/gin"
	mysql "index_Demo/dao/mysql"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/utils/middleware/auth"
)

// IsAdmin returns true if the user is an admin
func IsAdmin(ctx *gin.Context) bool {
	user := auth.CurrentUser(ctx)
	u := dal.User
	userInfo, _ := u.WithContext(ctx).Where(u.UID.Eq(user.UID)).First()
	return userInfo.IsAdmin
}

type Pagination struct {
	Total    int64 `json:"total"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
}

// QueryUsers returns a list of users and pagination information
func QueryUsers(ctx *gin.Context) ([]model.User, Pagination, error) {
	var users []model.User
	db := mysql.DB.GetDb()
	pagination := GetPagination(ctx)

	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	db.Model(&users).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order("create_at desc").Find(&users)
	if err := db.Error; err != nil {
		return nil, Pagination{}, nil
	}

	return users, pagination, nil
}

// QueryFdBack returns a list of feedback and pagination information
func QueryFdBack(ctx *gin.Context) ([]model.Feedback, Pagination, error) {
	var fdBack []model.Feedback
	db := mysql.DB.GetDb()
	pagination := GetPagination(ctx)

	offsetVal := (pagination.PageNum - 1) * pagination.PageSize
	db.Model(&fdBack).Count(&pagination.Total).Limit(pagination.PageSize).Offset(offsetVal).Order("create_at desc").Find(&fdBack)
	if err := db.Error; err != nil {
		return nil, Pagination{}, nil
	}

	return fdBack, pagination, nil
}

func GetPagination(ctx *gin.Context) Pagination {
	pagination := Pagination{}
	if err := ctx.BindJSON(&pagination); err != nil {
		// 返回默认值
		pagination.PageNum = 1
		pagination.PageSize = 10
	}
	if pagination.PageNum == 0 {
		pagination.PageNum = 1
	}
	if pagination.PageSize > 100 {
		pagination.PageSize = 100
	}
	if pagination.PageSize <= 0 {
		pagination.PageSize = 10
	}
	return pagination
}
