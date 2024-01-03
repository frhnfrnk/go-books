// book.go

package models

import (
	"github.com/frhnfrnk/go-books/utils"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null" validate:"required,min=3,max=255"`
	AuthorID    uint   `json:"author_id" gorm:"not null"`
	PublisherID uint   `json:"publisher_id" gorm:"not null"`
}

func (b *Book) Validate() error {
	return utils.ValidateStruct(b)
}
