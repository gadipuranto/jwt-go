package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserId      uint
	User        *User
}

func (p *Product) BeforeCreate(db *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		return errCreate
	}
	return nil
}

func (p *Product) BeforeUpdate(db *gorm.DB) error {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		return errUpdate
	}
	return nil
}
