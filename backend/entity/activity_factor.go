
package entity

import (
	"gorm.io/gorm"

)

type ActivityFactor struct {

	gorm.Model

	ActivityLevel string  //ระดับการออกกำลังกาย เช่น ไม่มีการออกกำลังกาย, กิจกรรมต่ำ, กิจกรรมเบา, กิจกรรมปานกลาง, กิจกรรมสูง, กิจกรรมสูงมาก
	EstimatedCalories int     //ปริมาณแคลอรี่ที่คาดว่าจะออกกำลังกาย, ค่าพลังงานที่ใช้ประมาณต่อวัน (kcal)

}