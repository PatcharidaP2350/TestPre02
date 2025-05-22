package config

import (

	"fmt"
	"time"
    "backend/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB


func DB() *gorm.DB {

   return db

}

func ConnectionDB() {

   database, err := gorm.Open(sqlite.Open("precap.db?cache=shared"), &gorm.Config{}) 

   if err != nil {

       panic("failed to connect database")

   }

   fmt.Println("connected database")

   db = database

}

func SetupDatabase() {


   db.AutoMigrate(

       &entity.Users{},

       &entity.Genders{},

	   &entity.ActivityFactor{},

	   &entity.ExerciseType{},

	   &entity.Exercise{},

	   &entity.ExerciseActivity{},

   )

   GenderMale := entity.Genders{GenderName: "Male"}

   GenderFemale := entity.Genders{GenderName: "Female"}


   db.FirstOrCreate(&GenderMale, &entity.Genders{GenderName: "Male"})

   db.FirstOrCreate(&GenderFemale, &entity.Genders{GenderName: "Female"})

   hashedPassword, _ := HashPassword("123456")

   BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")

   User := &entity.Users{

       FirstName: "ส้มจี๊ด",

       LastName:  "หวานเจี๊ยบ",

       Email:     "somsom@gmail.com",

       Age:       25,

       Password:  hashedPassword,

       BirthDate:  BirthDay,

       GenderID:  1,

	   Weight:    60.3,

	   Height:    170,


   }
   db.FirstOrCreate(User, &entity.Users{

       Email: "thrve@gmail.com",

   })


sedentaryactivityFactor := entity.ActivityFactor{ActivityLevel: "ไม่มีการออกกำลังกาย"}
lightActivityFactor := entity.ActivityFactor{ActivityLevel: "กิจกรรมต่ำ"}
moderateActivityactivityFactor := entity.ActivityFactor{ActivityLevel: "กิจกรรมเบา"}
activeactivityFactor := entity.ActivityFactor{ActivityLevel: "กิจกรรมปานกลาง"}
veryActiveactivityFactor := entity.ActivityFactor{ActivityLevel: "กิจกรรมสูง"}
extraActiveactivityFactor := entity.ActivityFactor{ActivityLevel: "กิจกรรมสูงมาก"}

db.FirstOrCreate(&sedentaryactivityFactor, &entity.ActivityFactor{ActivityLevel: "ไม่มีการออกกำลังกาย"})
db.FirstOrCreate(&lightActivityFactor, &entity.ActivityFactor{ActivityLevel: "กิจกรรมต่ำ"})
db.FirstOrCreate(&moderateActivityactivityFactor, &entity.ActivityFactor{ActivityLevel: "กิจกรรมเบา"})
db.FirstOrCreate(&activeactivityFactor, &entity.ActivityFactor{ActivityLevel: "กิจกรรมปานกลาง"})
db.FirstOrCreate(&veryActiveactivityFactor, &entity.ActivityFactor{ActivityLevel: "กิจกรรมสูง"})
db.FirstOrCreate(&extraActiveactivityFactor, &entity.ActivityFactor{ActivityLevel: "กิจกรรมสูงมาก"})
  
ExerciseType0 := entity.ExerciseType{TypeName: "Cardio"}
ExerciseType1 := entity.ExerciseType{TypeName: "Flex"}
ExerciseType2 := entity.ExerciseType{TypeName: "Strength"}

db.FirstOrCreate(&ExerciseType0, &entity.ExerciseType{TypeName: "Cardio"})
db.FirstOrCreate(&ExerciseType1, &entity.ExerciseType{TypeName: "Flex"})
db.FirstOrCreate(&ExerciseType2, &entity.ExerciseType{TypeName: "Strength"})

exercise := &entity.Exercise{
    ExerciseName:           "Running",
    ExerciseTypeId:         1,    // สมมุติว่า 1 = Cardio
    CaloriesBurndPerMinute: 10,
}

db.FirstOrCreate(exercise, &entity.Exercise{
    ExerciseName: "Running",
})

// สมมุติว่าคุณมี User ที่ ID = 1 และ Exercise ที่ ID = 1 อยู่แล้วในฐานข้อมูล

activityDate, _ := time.Parse("2006-01-02", "2025-05-20")

exerciseActivity := &entity.ExerciseActivity{
    ActivityName:     "Morning Run",
    UserID:           1, // รหัสของผู้ใช้งานที่ออกกำลังกาย
    ExerciseID:       1, // รหัสของประเภทการออกกำลังกาย เช่น Running
    Date:             activityDate,
    Duration:         30, // นาที
    CaloriesBurnd:    10 * 30, // สมมุติ Exercise ID 1 เบิร์น 10 cal/min
}

db.FirstOrCreate(exerciseActivity, &entity.ExerciseActivity{
    ActivityName: "Morning Run",
    UserID:       1,
    Date:         activityDate,
})


}