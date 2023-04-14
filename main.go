package main

import "github.com/gin-gonic/gin"



func main(){
	app := gin.Default()

	app.GET("/", index)
	app.GET("/contact",contact)
	app.GET("/about",about)

	app.Run()

}


func index(c *gin.Context){
	c.JSON(200,gin.H{
		"index":"Index page",
	})
}

func contact(c *gin.Context){
	c.JSON(200,gin.H{
		"contact":"contact_page",
	})
}


func about(c *gin.Context){
	c.JSON(200,gin.H{
		"about":"about page",
	})
}