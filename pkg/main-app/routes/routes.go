package routes

import (
	controller "github.com/Manas8803/DigiAuth/main-app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.RouterGroup) {
	router.POST("/wallet", controller.CreateWallet)
}
