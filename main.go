package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizzalbayu/test-todolist/controllers/activitycontroller"
	"github.com/rizzalbayu/test-todolist/controllers/todocontroller"
	"github.com/rizzalbayu/test-todolist/models"
)

func main() {
	route := gin.Default()
	models.ConnectDatabase()

	route.GET("/activity-groups", activitycontroller.Index)
	route.GET("/activity-groups/:id", activitycontroller.GetOne)
	route.POST("/activity-groups/", activitycontroller.Create)
	route.PATCH("/activity-groups/:id", activitycontroller.Update)
	route.DELETE("/activity-groups/:id", activitycontroller.Delete)
	route.GET("/todo-items", todocontroller.Index)
	route.GET("/todo-items/:id", todocontroller.GetOne)
	route.POST("/todo-items/", todocontroller.Create)
	route.PATCH("/todo-items/:id", todocontroller.Update)
	route.DELETE("/todo-items/:id", todocontroller.Delete)

	route.Run(":3000")
}
