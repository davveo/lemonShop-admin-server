package router

import (
	"github.com/davveo/lemonShop-admin-server/app/ctrs"
	"github.com/gin-gonic/gin"
)

func Init(route *gin.Engine) {
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	publicGroup := route.Group("/api/v1")

	checkCTRS := ctrs.NewCheckerController()
	{
		publicGroup.GET("check", checkCTRS.Check)
	}

	adminGroup := route.Group("/admin/admin")
	//adminGroup.Use(middleware.JWTAuth())
	{
		AdminRouterGroup(adminGroup)
	}
}
