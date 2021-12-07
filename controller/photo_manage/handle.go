package photo_manage

import (

	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)



func SetPhoto(ctx *gin.Context){

	fmt.Println("hello set photo")

	params := ctx.DefaultQuery("path","xxx")



	ctx.JSON(http.StatusOK, gin.H{
		"msg":params,
	})

}