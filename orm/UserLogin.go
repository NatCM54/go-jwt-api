package orm

import (
	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}
