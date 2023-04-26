package bing_wallpaper

import (
	"encoding/json"
	logs "github.com/danbai225/go-logs"
	"github.com/gin-gonic/gin"
	"index_Demo/gen/response"
	"io"
	"net/http"
)

type Article struct {
}

// Wallpaper
// @Summary	获取bing壁纸
// @Produce	json
// @Param		id	/wallpaper		int		true	"1"
// @Success	200	{object}	Article	"成功"
// @Failure	400	{object}	string	"获取壁纸失败"
// @Failure	500	{object}	string	"内部错误"
// @Router		/api/v1/articles/{id} [get]
func Wallpaper(ctx *gin.Context) {
	wallpaper, err := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "获取壁纸失败",
		})
		return
	}
	defer func(Body io.ReadCloser) {
		errs := Body.Close()
		if errs != nil {
			logs.Err(errs)
		}
	}(wallpaper.Body)

	var data map[string]interface{}
	if errs := json.NewDecoder(wallpaper.Body).Decode(&data); errs != nil {
		ctx.JSON(http.StatusInternalServerError, response.New("壁纸获取失败", errs))
		return
	}

	bing := "https://www.bing.com"
	url := bing + data["images"].([]interface{})[0].(map[string]interface{})["url"].(string)

	ctx.JSON(http.StatusOK, response.New(url, nil))
}
