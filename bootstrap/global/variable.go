package global

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
	DB          *gorm.DB
	Redis       *redis.Client
	Engine      *gin.Engine
)
