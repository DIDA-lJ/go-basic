package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	// gin 框架分组路由测试
	hello := r.Group("/hello")
	{
		hello.GET("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"name": "hello,world,gin!(Get)",
			})
		})
		hello.POST("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"name": "hello,world,gin!(Post)",
			})
		})
		hello.DELETE("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"name": "hello,world,gin!(Delete)",
			})
		})
		hello.PUT("", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"name": "hello,world,gin!(Put)",
			})
		})
	}

	err := r.Run(":9090")
	if err != nil {
		log.Println(err)
	}
}
