package server

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/chandrartw/cobacrud/controller"
)

func Routes(router *gin.Engine, inDB *controller.InDB){
		router.POST("/users/create", inDB.CreateUser)
		router.GET("/get/:id", inDB.GetUser)
		router.GET("/get-all", inDB.GetAllUser)
		router.PUT("/update", inDB.UpdateUser)
		router.DELETE("/delete/:id", inDB.DeleteUser)
}