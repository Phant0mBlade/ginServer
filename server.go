/*
This is where my starting code will be,
Code structure is still in development so things may change
*/

package main

import "github.com/gin-gonic/gin"

func ginRouter() *gin.Engine {
	return gin.Default()
}

// example server in gin
func server() {
	r := ginRouter()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func main() {
	server()
}
