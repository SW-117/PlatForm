package router


import(
	// "PlatForm/controller"
	"net/http"
	"github.com/gin-gonic/gin"
	"PlatForm/controller/photo_manage"
)


func helloHandle(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg":"hello www.baidu.com",
	})
}

func InitRouter() *gin.Engine{
	router := gin.Default()

	//首页路由
	router.GET("/hello",helloHandle)

	//controller路由
	controlGr := router.Group("/controller")

	photoGr := controlGr.Group("/photo_manage")

	photoGr.GET("/get",photo_manage.SetPhoto)

	return router

}