package main

import "github.com/gin-gonic/gin"


func main(){
	app := gin.Default()

	app.GET("/", index)
	app.GET("/about",about)
	app.GET("/contact",contact)

	app.Run()
}


func index(c *gin.Context){
	c.JSON(200,gin.H{
		"index": "index page",
	})
}

func about(c *gin.Context){
	c.JSON(200,gin.H{
		"about":"About page",
	})
}


func contact(c *gin.Context){
	c.JSON(200,gin.H{
		"contact": "Contact page",
	})
}