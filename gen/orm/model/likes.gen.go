// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLike = "likes"

// Like mapped from table <likes>
type Like struct {
	ID     int32 `gorm:"column:ID;primaryKey;autoIncrement:true" json:"ID"`
	UserID int32 `gorm:"column:UserID;primaryKey" json:"UserID"`
	PostID int32 `gorm:"column:PostID;primaryKey" json:"PostID"`
}

// TableName Like's table name
func (*Like) TableName() string {
	return TableNameLike
}
