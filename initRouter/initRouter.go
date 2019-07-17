package initRouter

import (
	"blog_api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	/**
	// 添加 Get 请求路由
	router.GET("/", retHelloGinAndMethod)
	// 添加 Post 请求路由
	router.POST("/", retHelloGinAndMethod)
	// 添加 Put 请求路由
	router.PUT("/", retHelloGinAndMethod)
	// 添加 Delete 请求路由
	router.DELETE("/", retHelloGinAndMethod)
	// 添加 Patch 请求路由
	router.PATCH("/", retHelloGinAndMethod)
	// 添加 Head 请求路由
	router.HEAD("/", retHelloGinAndMethod)
	// 添加 Options 请求路由
	router.OPTIONS("/", retHelloGinAndMethod)

	router.GET("/user/:name",handler.UserSave)
	router.GET("/user", handler.UserSaveByQuery)

	**/

	/**
	路由分组
	*/
	router.LoadHTMLGlob("./templates/*")
	router.Static("/statics", "./statics")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	index := router.Group("/")
	{
		index.GET("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}
	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	//router.LoadHTMLGlob("templates/*")
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
