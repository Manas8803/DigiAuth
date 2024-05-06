package controllers

import (
	"log"
	"net/http"

	"github.com/Manas8803/DigiAuth/pkg/lib/ledger"
	"github.com/Manas8803/DigiAuth/pkg/main-app/responses"
	"github.com/gin-gonic/gin"
)

func CreateWallet(r *gin.Context) {

	//! MAKE CHANGES HERE

	r.JSON(http.StatusOK, responses.CreateWalletResponse{Message: "Successfully created Wallet", Data: map[string]interface{}{"Some Data":"Data1"}}) //! CHANGE THIS
}
func GenerateDID(r *gin.Context) {

	res, err := ledger.RegisterDID("firstuser", "firstuserseed", "STEWARD")
	if err != nil {
		log.Println("Error in registering DID", err)
		return
	}

	r.JSON(http.StatusOK, responses.CreateWalletResponse{Message: "Successfully created Wallet", Data: map[string]interface{}{"Data":res}})
}