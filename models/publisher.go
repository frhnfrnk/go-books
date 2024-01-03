// publisher.go

package models

import (
	"github.com/frhnfrnk/go-books/utils"
	"github.com/jinzhu/gorm"
)

type Publisher struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique;not null" validate:"required,min=3,max=100"`
	Books []Book `json:"books"`
}

func (p *Publisher) Validate() error {
	return utils.ValidateStruct(p)
}
