package models

import (
	"github.com/frhnfrnk/go-books/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=4,max=20"`
	Password string `json:"password,omitempty" validate:"required,min=6"`
}

func (u *User) Validate() error {
	return utils.ValidateStruct(u)
}
