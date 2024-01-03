// author.go

package models

import (
	"github.com/frhnfrnk/go-books/utils"
	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null" validate:"required,min=3,max=100"`
	Books []Book `json:"books"`
}

func (a *Author) Validate() error {
	return utils.ValidateStruct(a)
}
