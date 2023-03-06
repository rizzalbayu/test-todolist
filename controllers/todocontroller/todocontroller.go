package todocontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rizzalbayu/test-todolist/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var todos []models.Todo

	activityId := c.DefaultQuery("activity_group_id", "")
	if activityId != "" {
		models.DB.Where("activity_group_id = ?", activityId).Find(&todos)
	} else {
		models.DB.Find(&todos)
	}
	// models.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func GetOne(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := models.DB.First(&todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Todo with ID " + id + " Not Found", "data": gin.H{}})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todo})
}

func Create(c *gin.Context) {
	var todo models.Todo
	var activity models.Activity

	if err := c.ShouldBindJSON(&todo); err != nil {
		var emailErr, activityErr error
		for _, v := range err.(validator.ValidationErrors) {
			if v.Field() == "Title" {
				emailErr = v
				break
			}
			if v.Field() == "ActivityGroupId" {
				activityErr = v
				break
			}
		}
		if emailErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null", "data": gin.H{}})
			return
		}
		if activityErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "activity_group_id cannot be null", "data": gin.H{}})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
		return
	}
	if err := models.DB.First(&activity, todo.ActivityGroupId).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "Activity with ID "+ fmt.Sprintf("%v",todo.ActivityGroupId) +" Not Found", "data": gin.H{}})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
			return
		}
	}

	models.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": "Success", "message": "Success", "data": todo})

}
func Update(c *gin.Context) {
	var todo models.Todo
	// var activity models.Activity
	id := c.Param("id")
	
	if err := models.DB.First(&todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Todo with ID " + id + " Not Found", "data": gin.H{}})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
			return
		}
	}

	var todoUpdate models.TodoUpdate
  if err := c.ShouldBindJSON(&todoUpdate); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
  }

	if todoUpdate.Title != "" {todo.Title = todoUpdate.Title}
	if todoUpdate.IsActive != nil {todo.IsActive = *todoUpdate.IsActive}
	// if todoUpdate.Priority != "" {todo.Priority = todoUpdate.Priority}
	// if todoUpdate.ActivityGroupId != "" {todo.ActivityGroupId = todoUpdate.ActivityGroupId}

	// if err := models.DB.First(&activity, todo.ActivityGroupId).Error; err != nil {
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "Activity with ID "+ fmt.Sprintf("%v",todo.ActivityGroupId) +" Not Found", "data": gin.H{}})
	// 		return
	// 	default:
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
	// 		return
	// 	}
	// }

	// if models.DB.Model(&todo).Where("id = ?", id).Updates(&todo).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Todo with ID " + id + " Not Found", "data": gin.H{}})
	// 	return
	// }

	// Save changes to database
	if err := models.DB.Save(&todo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Error saving todo",
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todo})
}

func Delete(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if models.DB.Delete(&todo, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Todo with ID " + id + " Not Found", "data": gin.H{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": gin.H{}})
}
