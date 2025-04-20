package router

import (
	"file-manager/controller"
	"file-manager/middleware"
	"file-manager/svc"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine, ctx *svc.ServiceContext) {
	routeMiddleware := middleware.NewShortRouteMiddleware(ctx)
	r := g.Group("/api/v1")
	fileOperate := controller.NewFileOperator(ctx)
	r.POST("/upload", fileOperate.UploadFile)
	r.GET("/resource/:filepath", fileOperate.GetResourceID)
	g.NoRoute(routeMiddleware.Handle())
}
