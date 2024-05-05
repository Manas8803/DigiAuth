package controllers

import (
	"net/http"

	"github.com/Manas8803/DigiAuth/pkg/main-app/responses"
	"github.com/gin-gonic/gin"
)

func CreateWallet(r *gin.Context) {

	//! MAKE CHANGES HERE

	r.JSON(http.StatusOK, responses.CreateWalletResponse{Message: "Successfully created Wallet", Data: map[string]interface{}{"Some Data":"Data1"}}) //! CHANGE THIS
}