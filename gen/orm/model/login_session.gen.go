// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLoginSession = "login_session"

// LoginSession mapped from table <login_session>
type LoginSession struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // token_Id
	Token     string    `gorm:"column:token;not null" json:"token"`                // 令牌
	UID       int32     `gorm:"column:uid;not null" json:"uid"`                    // 用户ID
	LoginTime time.Time `gorm:"column:login_time" json:"login_time"`               // 登陆时间
	LoginIP   *string   `gorm:"column:login_ip" json:"login_ip"`                   // 登陆IP
}

// TableName LoginSession's table name
func (*LoginSession) TableName() string {
	return TableNameLoginSession
}