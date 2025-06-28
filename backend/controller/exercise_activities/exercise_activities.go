package exercise_activities

import (
	"fmt"
	"net/http"
	"time"

	"backend/config"

	"backend/entity"

	"github.com/gin-gonic/gin"
)


func GetExerciseActivitiesbyID(c *gin.Context) {


   ID := c.Param("id")

   var exercise_activity entity.ExerciseActivity


   db := config.DB()

   results := db.Preload("Exercise").First(&exercise_activity, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if exercise_activity.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, exercise_activity)


}


func CreateExerciseActivity (c *gin.Context) {

	fmt.Println("Creating or Updating Exercise Activity")

	db := config.DB()

	var input struct {
		ActivityName  string    `json:"activity_name"`           // ชื่อกิจกรรมการออกกำลังกาย
    	UserID        uint      `json:"user_id"`
    	Date          time.Time `json:"date"`                    // วันที่
    	Duration      int       `json:"duration"`                // เวลา (นาที) ที่ User ออกกำลังกาย
    	CaloriesBurnd int       `json:"calories_burnd"`           // แคลอรีที่เผา
		UserId 		  uint
		ExerciseID uint
	}

	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// สร้างหรืออัปเดต ExerciseActivity ในฐานข้อมูล
	exerciseActivity := entity.ExerciseActivity{
		ActivityName:  input.ActivityName,
		UserID:        input.UserID,
		Date:          input.Date,
		Duration:      input.Duration,
		CaloriesBurnd: input.CaloriesBurnd,
		UserId: 		  input.UserId,
		ExerciseID:    input.ExerciseID,
	}

	// เริ่มต้นการเชื่อมต่อฐานข้อมูล
	if err := db.Create(&exerciseActivity).Error; err != nil {
		fmt.Println("Error saving medical record:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save medical record", "details": err.Error()})
		return
	}

	// ส่งการตอบกลับที่ประสบความสำเร็จ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Medical record created and TakeAHistory handled successfully",
		"data":    exerciseActivity,
	})
}

func UpdateExerciseActivitybyID(c *gin.Context) {

	fmt.Print("Updating Exercise Activity")
	db := config.DB()

	// รับ ExerciseActivityID จากพารามิเตอร์  URL
	id := c.Param("id")
	var activity entity.ExerciseActivity
	if err := db.First(&activity, id).Error; err != nil {
		fmt.Println("Record not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "TakeAHistory record not found"})
		return
	}

	// กำหนดโครงสร้างข้อมูลที่ได้รับจาาก Request
	var input struct {
		ActivityName  string    `json:"activity_name"`           // ชื่อกิจกรรมการออกกำลังกาย
    	UserID        uint      `json:"user_id"`
    	Date          time.Time `json:"date"`                    // วันที่
    	Duration      int       `json:"duration"`                // เวลา (นาที) ที่ User ออกกำลังกาย
    	CaloriesBurnd int       `json:"calories_burnd"`           // แคลอรีที่เผา
		UserId 		  uint
		ExerciseID uint
	}

	// ดึงข้อมูล JSON จากคำขอ (Request) และตรวจสอบความถูกต้อง
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data", "details": err.Error()})
		return
	}

	// อัปเดต Fields ที่ส่งมาใน Request
	if input.ActivityName != "" {
		activity.ActivityName = input.ActivityName
	}

	
}