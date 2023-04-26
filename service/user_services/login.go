package user_services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"index_Demo/app/request"
	"index_Demo/dao/redisServer"
	"index_Demo/gen/orm/dal"
	"index_Demo/gen/orm/model"
	"index_Demo/gen/response"
	"index_Demo/utils/serverUtils"
	"index_Demo/utils/text"
	"index_Demo/utils/validateUtils"
	"net/http"
	"time"
)

func Login(ctx *gin.Context) {
	loginReq := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginReq)
	loginIp := serverUtils.GetRealIP(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.New("绑定数据失败", err))
		return
	}
	message, hasErr := validateUtils.ReturnValidateMessage(&loginReq, err)
	if hasErr {
		ctx.JSON(http.StatusUnprocessableEntity, response.New(message, nil))
		return
	}

	u := dal.User

	userM, err := u.WithContext(ctx).Where(u.Username.Eq(loginReq.Username)).First()
	if userM == nil {
		ctx.JSON(http.StatusBadRequest, response.New("用户不存在", err.Error()))
		return
	}

	//获取数据库密码 并对比
	sqlUser, err := u.WithContext(ctx).Where(u.Username.Eq(loginReq.Username)).Select(u.UID, u.Username, u.Password).First()
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(sqlUser.Password), []byte(loginReq.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.New("密码错误", err.Error()))
		return
	}

	//生成token
	tokenM := model.LoginSession{
		Token:     text.GetUUID(),
		UID:       userM.UID,
		LoginTime: time.Now(),
		LoginIP:   &loginIp,
	}
	loginSession := dal.LoginSession

	//把登陆信息token存入mysql
	loginSession.WithContext(ctx).Create(&tokenM)

	jsonM, _ := json.Marshal(userM)

	//token存入redis
	err = redisServer.Set(fmt.Sprintf("user_token_%s", tokenM.Token), string(jsonM), time.Duration(viper.GetInt("redis.tokenTime"))*time.Minute)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("登录失败（token生成失败) 请联系管理员", err.Error()))
		return
	}
	redisServer.PutSet(fmt.Sprintf("user_tokens_%s", userM.Username), []string{tokenM.Token})

	ctx.JSON(http.StatusOK, response.New("登陆成功", map[string]interface{}{
		"loginIp":   loginIp,
		"loginTime": time.Now(),
		"user": gin.H{
			"username": userM.Username,
			"token":    tokenM.Token,
		},
		//"user":      userM,
	}))
}
