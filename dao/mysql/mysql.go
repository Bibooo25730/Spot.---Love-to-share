package nosql

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"index_Demo/gen/orm/dal"
	"time"
)

// Predicate is a string that acts as a condition in the where clause
type Predicate string

var _ Repo = (*dbRepo)(nil)

type Repo interface {
	i()
	GetDb() *gorm.DB
	DbClose() error
}

type dbRepo struct {
	Db *gorm.DB
}

var DB *dbRepo

func Init() error {

	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetUint16("mysql.port")
	dbName := viper.GetString("mysql.db")

	db, err := dbConnect(
		user,
		password,
		fmt.Sprintf("%s:%d", host, port), dbName,
	)
	if err != nil {
		return err
	}
	DB = &dbRepo{Db: db}
	return nil
}

func (d *dbRepo) i() {}

func (d *dbRepo) GetDb() *gorm.DB {
	return d.Db.Debug()
}

func (d *dbRepo) DbClose() error {
	sqlDB, err := d.Db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbName,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	dal.SetDefault(db)

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(100)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(10)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * 10)

	return db, nil
}
