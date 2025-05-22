

package entity

import (
	"gorm.io/gorm"
)

type ExerciseType struct {

	gorm.Model

	TypeName string

}



