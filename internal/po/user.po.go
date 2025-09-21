package po

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement; comment:'primary key id'"`
	UserName string `gorm:"column:user_name"`
	IsActive bool   `gorm:"column:is_active; type: boolean"`
	Roles    []Role `gorm:"many2many:go_user_roles"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
