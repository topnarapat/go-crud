package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/topnarapat/go-crud/configs"
	"github.com/topnarapat/go-crud/models"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	configs.DB.Raw("select * from users order by id desc").Scan(&users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  input.Password,
		IsAdmin:   input.IsAdmin,
	}

	// check duplicate email
	userExist := configs.DB.Where("email = ?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exist"})
		return
	}

	result := configs.DB.Create(&user)

	// db error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": "Registration Successful",
	})
}
