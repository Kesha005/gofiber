package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{
	ID string	`json:"id"` 
	Name string		`json:"name"`
	Surname string 	`json:"surname"`
	Age int			`json:"age"`
}



var user = []User{
	User{ID:"2" , Name: "Kerimberdi" ,Surname: "Saparow", Age: 17 },
	User{ID:"3" , Name: "Kakamyrat" ,Surname: "Saparow", Age: 14 },
}

func main(){

	app := gin.Default()
	app.GET("/",getAllUsers)
	app.GET("user/:id",getUserById)
	app.POST("/post_user",newUser)

	app.Run(":8080")
}


func getAllUsers(c *gin.Context){
	c.JSON(http.StatusOK,user)
}


func getUserById(c *gin.Context){
	for _ ,finded := range user{
		if finded.ID == c.Param("id"){
			c.JSON(http.StatusFound,finded)
			return
		}
		
	}
	c.JSON(http.StatusNotFound,gin.H{
		"message":"User not foud",
	})
}


func newUser(c *gin.Context){
	var new_user User

	if err := c.BindJSON(&new_user); err!= nil{
		c.JSON(http.StatusFailedDependency,err.Error())
		return 
	}

	user = append(user, new_user)
	c.JSON(http.StatusCreated,new_user)
}