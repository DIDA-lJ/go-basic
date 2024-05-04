package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Id      int64  `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address"`
}

type Group struct {
	Groups []string `form:"group"`
}

type MapTest struct {
	AddressMap map[string]string `form:"addressMap"`
}

type LoginUser struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	CheckCode string `json:"checkcode"`
}

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
	// Get 获取普通参数,获取方法 1，直接获取
	r.GET("/user/save", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		// 地址不存在，给一个默认值
		address := ctx.DefaultQuery("address", "beijing")
		ctx.JSON(200, gin.H{
			"code":    200,
			"id":      id,
			"name":    name,
			"address": address,
		})
	})

	// Get 获取普通参数,获取方法 2 ，对象获取
	r.GET("/user/test", func(ctx *gin.Context) {
		var user User
		// BindQuery 返回的是一个列表参数
		err := ctx.ShouldBind(&user)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, user)
	})

	// Get 数组参数获取
	r.GET("/user/array", func(ctx *gin.Context) {
		var group Group
		err := ctx.ShouldBind(&group)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, group)
	})
	// Get map 类型参数
	r.GET("/user/map", func(ctx *gin.Context) {
		var mapTest MapTest
		err := ctx.ShouldBind(&mapTest)
		if err != nil {
			log.Println(err)
		}
		mapTest.AddressMap = ctx.QueryMap("addressMap")
		ctx.JSON(200, mapTest)
	})

	// Post 获取参数 （表单）
	r.POST("/user/login", func(ctx *gin.Context) {
		var loginUser LoginUser
		err := ctx.ShouldBind(&loginUser)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, gin.H{"code": "200", "account": loginUser.Account, "password": loginUser.Password})
	})

	// Post 获取参数 （Json）
	r.POST("/user/loginJson", func(ctx *gin.Context) {
		var loginUser LoginUser
		err := ctx.ShouldBindJSON(&loginUser)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(200, gin.H{"code": "200", "account": loginUser.Account, "password": loginUser.Password, "checkcode": loginUser.CheckCode})
	})

	// Post 路径参数(路径参数测试)
	r.POST("/user/urlTest/:id/:name", func(ctx *gin.Context) {
		id := ctx.Param("id")
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{"id": id, "name": name})
	})

	// Post 文件参数获取,文件上传测试
	r.POST("/user/fileTest", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			log.Println(err)
		}
		value := form.Value
		files := form.File
		for _, fileArray := range files {
			for _, v := range fileArray {
				ctx.SaveUploadedFile(v, "./"+v.Filename)
			}
		}
		ctx.JSON(200, value)
	})

	err := r.Run(":9090")
	if err != nil {
		log.Println(err)
	}

}
