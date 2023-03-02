package todocontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rizzalbayu/test-todolist/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var todos []models.Todo

	models.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func GetOne(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := models.DB.First(&todo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Activity with ID " + id + " Not Found", "data": gin.H{}})
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

	if err := c.ShouldBindJSON(&todo); err != nil {
		var emailErr error
		for _, v := range err.(validator.ValidationErrors) {
			if v.Field() == "Title" {
				emailErr = v
				break
			}
		}
		if emailErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null", "data": gin.H{}})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
		return
	}

	models.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todo})

}
func Update(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
		return
	}

	if models.DB.Model(&todo).Where("id = ?", id).Updates(&todo).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "todo with ID " + id + " Not Found", "data": gin.H{}})
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
