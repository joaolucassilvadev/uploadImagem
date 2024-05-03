package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/imageupload/controller"
)

func Routes(eng *gin.Engine) {
	eng.POST("upload", controller.UploadImg)
	eng.GET("/thumbnail", controller.Thumbnail)
	eng.GET("/image", controller.Image)
	eng.GET("/", controller.Html)
}
