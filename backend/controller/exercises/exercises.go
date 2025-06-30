package exercises


import (

   "net/http"

   "backend/config"

   "backend/entity"

   "github.com/gin-gonic/gin"

)


func ListExercises(c *gin.Context) {


   db := config.DB()


   var exercises []entity.Exercise

   db.Find(&exercises)


   c.JSON(http.StatusOK, &exercises)


}