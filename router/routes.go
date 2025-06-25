package routes

import (
	"github.com/Triptiverma003/go-gin-auth/controller"
	"github.com/Triptiverma003/go-gin-auth/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine	) {
	r.POST("/login" , controller.Login)
	r.POST("/register" , controller.Register)
	// r.GET("/logout" , controller.LogOut)

	private := r.Group("/private")

	private.Use(middleware.Authenticate)
	
	private.GET("/refreshtoken" , controller.RefreshToken)
	// r.GET("/logout" , controller.RefreshToken)
}