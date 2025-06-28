

package entity

import (
	"gorm.io/gorm"
)

type ExerciseType struct {

	gorm.Model

	TypeName 	string 	`json:"type_name"`

}



