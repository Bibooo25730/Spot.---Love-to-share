package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type config struct {
	Port  uint16 `json:"port"`
	Mysql struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Port     uint16 `json:"port"`
		DB       string `json:"db"`
	} `json:"mysql"`
	CollectionAddress string `json:"collection_address"`
	Jwt               struct {
		TokenExpire int `json:"tokenExpire"`
	}
}

var Config = config{}

// LoadConfig 使用os.ReadFile读取配置文件
func LoadConfig() error {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		return err
	}

	return nil
}

// InfoConfig 使用viper读取配置文件
func InfoConfig() error {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Load Config Error: %s", err.Error()))
	}
	return err
}
