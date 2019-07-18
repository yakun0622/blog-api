package initRouter

import (
	"blog_api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

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
	btcUsdRouter := router.Group("btc_usd")
	{
		btcUsdRouter.GET("/list", handler.QueryHlocList)
	}
	return router
}
