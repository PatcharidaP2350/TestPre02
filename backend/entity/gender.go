package entity

import (
	"gorm.io/gorm"
)

type Genders struct {

	gorm.Model

	GenderName string
	
}
