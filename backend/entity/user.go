package entity

import (
	"time"
	"gorm.io/gorm"
)

type Users struct {

	gorm.Model

	FirstName      string
	LastName       string
	Age            uint8
	BirthDate      time.Time
	PhoneNumber    string
	Weight         float32
	Height         float32
	Email     string	
	Password  string
	GenderID  uint      `json:"gender_id"`
    Gender    *Genders  `gorm:"foreignKey: gender_id" json:"gender"`
	FactorID  uint      `json:"factor_id"`
	Factor    *ActivityFactor  `gorm:"foreignKey: factor_id" json:"factor"`

}