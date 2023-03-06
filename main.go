package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizzalbayu/test-todolist/controllers/activitycontroller"
	"github.com/rizzalbayu/test-todolist/controllers/todocontroller"
	"github.com/rizzalbayu/test-todolist/models"
)

func main() {
	// router := gin.Default()
	router := gin.New()
	router.Use(
			gin.LoggerWithWriter(gin.DefaultWriter, "/pathsNotToLog/"),
			gin.Recovery(),
	)
	models.ConnectDatabase()

	router.GET("/activity-groups", activitycontroller.Index)
	router.GET("/activity-groups/:id", activitycontroller.GetOne)
	router.POST("/activity-groups", activitycontroller.Create)
	router.PATCH("/activity-groups/:id", activitycontroller.Update)
	router.DELETE("/activity-groups/:id", activitycontroller.Delete)
	router.GET("/todo-items", todocontroller.Index)
	router.GET("/todo-items/:id", todocontroller.GetOne)
	router.POST("/todo-items", todocontroller.Create)
	router.PATCH("/todo-items/:id", todocontroller.Update)
	router.DELETE("/todo-items/:id", todocontroller.Delete)

	router.Run(":3030")
}
