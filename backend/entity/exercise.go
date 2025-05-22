package entity

import (

	"gorm.io/gorm"
)

type Exercise struct {

	gorm.Model

	ExerciseName string
	ExerciseTypeId uint           //(Foreign Key to ExerciseType)
	ExerciseType *ExerciseType `gorm:"foreignKey: exercise_type_id" json:"exerciseType"`
	CaloriesBurndPerMinute int
	
}	