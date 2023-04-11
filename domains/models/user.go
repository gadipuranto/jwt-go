package models

type User struct {
	GormModel
	FullName string `gorm:"not_null" json:"full_name" valid:"required~Your full name is required"`
	Email    string `gorm:"not_null;uniqueIndex" json:"full_name" valid:"required~Your full name is required"`
}
