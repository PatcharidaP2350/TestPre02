package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/config"
	"backend/controller/exercises"
    "backend/controller/exercise_activities"
	"backend/controller/genders"
	"backend/controller/users"
	"backend/middlewares"
)


const PORT = "8000"


func main() {


   // open connection database

   config.ConnectionDB()


   // Generate databases

   config.SetupDatabase()


   r := gin.Default()


   r.Use(CORSMiddleware())


   // Auth Route

   r.POST("/signup", users.SignUp)

   r.POST("/signin", users.SignIn)


   router := r.Group("/")

   {

       router.Use(middlewares.Authorizes())


       // User Route

       router.PUT("/user/:id", users.Update)

       router.GET("/users", users.GetAll)

       router.GET("/user/:id", users.Get)

       router.DELETE("/user/:id", users.Delete)


   }

   // Gender Route
   r.GET("/genders", genders.GetAll)

   // Exercise Records Route
   r.GET("/exercises", exercises.ListExercises)
   

   // Exercise Activity Route
   r.GET("/exercise_activity/:id", exercise_activities.GetExerciseActivitiesbyID)  // Get
   r.GET("/exercise_activities/user/:user_id", exercise_activities.GetExerciseActivitiesbyUserID)

   r.POST("/exercise_activity", exercise_activities.CreateExerciseActivity)      // Create
   r.PUT("/exercise_activity/:id", exercise_activities.UpdateExerciseActivitybyID)   // Update
   r.DELETE("/exercise_activity/:id", exercise_activities.DeleteExerciseActivitybyID) // Delete

   r.GET("/", func(c *gin.Context) {

    c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)

   })


   // Run the server


   r.Run("localhost:" + PORT)


}


func CORSMiddleware() gin.HandlerFunc {

   return func(c *gin.Context) {

       c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

       c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

       c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

       c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")


       if c.Request.Method == "OPTIONS" {

           c.AbortWithStatus(204)

           return

       }


       c.Next()

   }

}