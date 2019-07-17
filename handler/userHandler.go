package handler

import (
	"blog_api/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	println(username)
	context.String(http.StatusOK, "用户已经保存")
}

// 通过 query 方法进行获取参数
func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "用户:"+username+",年龄:"+age+"已经保存")
}

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
		context.Redirect(http.StatusMovedPermanently, "/")
	}
}
