package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}



func main(){
	app := gin.Default()
	app.POST("/user",post)

	app.Run()

}


func post(c *gin.Context){
	var user Login
	if err := c.ShouldBindJSON(&user);err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"status":"all right",
	})
}