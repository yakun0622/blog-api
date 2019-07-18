package handler

import (
	"blog_api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryHlocList(context *gin.Context) {
	coin := model.Coin{}
	list, dataArray, _ := coin.QueryList("btc_usd_15m")
	println(list)

	context.JSON(http.StatusOK, gin.H{
		"btc_usd_list":  list,
		"btc_usd_array": dataArray,
	})
}
