package entity

import (

	"gorm.io/gorm"

	"time"
)

type ExerciseActivity struct {   //กิจกรรมที่ User ออก

	gorm.Model

	ActivityName string    //ชื่อกิจกรรมการออกกำลังกาย
	UserID uint
	Date time.Time          //วันที่
	Duration int  //เวลาที่ User ออกกำลังกาย
	CaloriesBurnd int   //(คำนวณ = calories_burned_per_minute * duration) แคลที่เผา
	UserId uint
	User *Users `gorm:"foreignKey: user_id" json:"user"`         //(Foreign Key to User)
	ExerciseID uint
	Exercise *Exercise `gorm:"foreignKey: exercise_id" json:"exercise"`             //(Foreign Key to Exercise)
	
}
