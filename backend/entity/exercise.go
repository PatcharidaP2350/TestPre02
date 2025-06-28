package entity

import (

	"gorm.io/gorm"
)

type Exercise struct {

	gorm.Model

	ExerciseName           string        `json:"exercise_name"`
    ExerciseTypeId         uint          `json:"exercise_type_id"`
    ExerciseType           *ExerciseType `gorm:"foreignKey:exercise_type_id" json:"exercise_type"`
    CaloriesBurndPerMinute int           `json:"calories_burnd_per_minute"`
	
}	