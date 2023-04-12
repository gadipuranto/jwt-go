package models

import (
	"github.com/asaskevich/govalidator"
	"go-jwt/helpers"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not_null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not_null;uniqueIndex" json:"email" form:"email" valid:"required~Your full email is required,email~invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~ Your password is required,minstringlength(6)~Password has to have min length of 6 charachters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u User) BeforeCreate(db *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		return errCreate
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}
