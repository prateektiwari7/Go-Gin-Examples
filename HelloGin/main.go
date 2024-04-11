package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/Hello", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
	})
	r.Run()
}

/*
Map object written
ctx.Json(200,gin.H{
	"message" : "Hello"
})
*/
