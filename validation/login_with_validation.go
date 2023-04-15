package main

import (
	"github.com/gin-gonic/gin"
)


type User struct{
	Name string		`json:"name" binding:"required"`
	Email string 	`json:"email" binding:"required"`
	Password string	`json:"password" binding:"required"`
}


var users = []User{
	User{"Kerim" , "kerim@gmail.com" , "12344321" },
	User{"Mak" , "mak@gmail.com" ,"12344321"},
} 


func main(){
	
	App := gin.Default()

	App.POST("login",login)

	App.Run(":80")
}

func login(c *gin.Context){
	var user User
	if err := c.ShouldBindJSON(&user);err != nil{
		c.JSON(400,gin.H{
			"message":"Unvalidated data",
			"status":false,
		})
		return 
	}
	user_email := c.Param("email")
	user_pass := c.Param("password")

	for _,user := range users {
		if user.Email == user_email && user.Password == user_pass {
			c.JSON(200,gin.H{
				"message":"User login ,Hello " + user.Name,
				"status":true,
			})
			return 
		}
		c.JSON(403,gin.H{
			"status":false,
			"message":"incorrect user name or email",
		})
		return 
	} 
}