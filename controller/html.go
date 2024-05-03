package controller

import "github.com/gin-gonic/gin"

func Html(ctx *gin.Context) {
	ctx.File("index.html")
}
