package activitycontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rizzalbayu/test-todolist/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var activities []models.Activity

	models.DB.Find(&activities)
	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activities})
}

func GetOne(c *gin.Context) {
	var activity models.Activity
	id := c.Param("id")

	if err := models.DB.First(&activity, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Activity with ID " + id + " Not Found", "data": gin.H{}})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activity})

}

func Create(c *gin.Context) {
	var activity models.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
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

	models.DB.Create(&activity)
	c.JSON(http.StatusCreated, gin.H{"status": "Success", "message": "Success", "data": activity})

}

func Update(c *gin.Context) {
	var activity models.Activity
	id := c.Param("id")

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
		return
	}

	if models.DB.Model(&activity).Where("id = ?", id).Updates(&activity).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Activity with ID " + id + " Not Found", "data": gin.H{}})
		return
	}

	if err := models.DB.First(&activity, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Activity with ID " + id + " Not Found", "data": gin.H{}})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Failed", "message": err.Error(), "data": gin.H{}})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": activity})
}

func Delete(c *gin.Context) {
	var activity models.Activity

	id := c.Param("id")
	fmt.Print(id)

	if models.DB.Delete(&activity, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "Not Found", "message": "Activity with ID " + id + " Not Found", "data": gin.H{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": gin.H{}})
}
