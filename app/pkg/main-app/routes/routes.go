package routes

import (
	controller "github.com/Manas8803/DigiAuth/pkg/main-app/controllers"

	"github.com/gin-gonic/gin"
)

func LedgerRoute(router *gin.RouterGroup) {
	router.POST("/wallet", controller.CreateWallet)
	router.POST("/certificate", controller.IssueCeritificate)
}
