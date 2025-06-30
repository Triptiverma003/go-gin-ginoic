package controller

import (
	"log"

	"github.com/Triptiverma003/go-gin-auth/database"
	"github.com/Triptiverma003/go-gin-auth/helper"
	"github.com/Triptiverma003/go-gin-auth/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type formData struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	returnObject := gin.H{
		"status" : "Ok",
		"msg" : "Login route",
	}

	var formData formData

	if err := c.ShouldBind(&formData); err != nil{
		log.Println("error in binding data")

		c.JSON(400, returnObject)
		return
	}
	var user model.User
	database.DBConn.First(&user , "email=?" , formData.Email)

	if user.ID == 0{
		returnObject["msg"] = "User Not found"

		c.JSON(400 , returnObject)
		return
	}

	//Validate Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password) , []byte(formData.Password))

	if err != nil {
		log.Println("Inalid Password")

		returnObject["msg"] = "Password Doesnt match"
		c.JSON(401 , returnObject)
		return
	}

	// create Token

	token, err:= helper.GenerateToken(user)

	if err != nil{
		returnObject["msg"] = "Token Generation error"
		c.JSON(401 , returnObject)
		return 
	}

	returnObject["token"] = token
	returnObject["user"] = user	
	returnObject["status"] = "Ok"
	returnObject["msg"] = "User authenticated"
	c.JSON(200 , returnObject)
}

func Register(c *gin.Context) {
	returnObject := gin.H{
		"status" : "Ok",
		"msg" : "Register route",
	}
	c.JSON(200 , returnObject)

	var formData formData

	if err := c.ShouldBind(&formData); err != nil {
		log.Println("error in JSON binding")
		returnObject["msg"] = "error in binding.."
		c.JSON(400 , returnObject)
		return
	}

	var user model.User

	user.Email = formData.Email
	user.Password = helper.HashPassword(formData.Password)


	result := database.DBConn.Create(&user)

	if result.Error != nil{
		log.Println(result.Error)
		returnObject["msg"] = "User Already Exist..."
		c.JSON(400 , returnObject)
		return
	}
	returnObject["msg"] = "User Added successfully"
	c.JSON(201 , returnObject)
}



func LogOut() {}

func RefreshToken(c *gin.Context) {
	returnObject := gin.H{
		"status" : "Ok",
		"msg" : "Refresh Token Route",
	}
	c.JSON(200 , returnObject)
}