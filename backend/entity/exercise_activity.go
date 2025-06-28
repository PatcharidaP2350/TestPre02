package entity

import (

	"gorm.io/gorm"

	"time"
)

type ExerciseActivity struct {   //กิจกรรมที่ User ออก

	gorm.Model

	ActivityName  string    `json:"activity_name"`           // ชื่อกิจกรรมการออกกำลังกาย
    UserID        uint      `json:"user_id"`
    Date          time.Time `json:"date"`                    // วันที่
    Duration      int       `json:"duration"`                // เวลา (นาที) ที่ User ออกกำลังกาย
    CaloriesBurnd int       `json:"calories_burnd"`           // แคลอรีที่เผา
	UserId 		  uint
	User          *Users     `gorm:"foreignKey:UserId"`         //(Foreign Key to User)
	ExerciseID uint
	Exercise      *Exercise  `gorm:"foreignKey:ExerciseID"`             //(Foreign Key to Exercise)
	
}
	