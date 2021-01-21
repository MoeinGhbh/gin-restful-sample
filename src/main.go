package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"messenger": "Hello World",
	})
}

func about(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "dfghdgfdhfg",
	})
}

func HomePostPage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"messenger": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func pathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})

}

func main() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.GET("/about", about)
	r.POST("/", HomePostPage)
	r.GET("/query", QueryStrings)             // /query?name=moein&age=39
	r.GET("/path/:name/:age", pathParameters) // /path/moein/54/

	r.Run()
}
