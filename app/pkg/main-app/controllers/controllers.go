package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	db "github.com/Manas8803/DigiAuth/db/sqlconfig"
	"github.com/Manas8803/DigiAuth/pkg/lib/configs"
	"github.com/Manas8803/DigiAuth/pkg/lib/ledger"
	"github.com/Manas8803/DigiAuth/pkg/main-app/models"
	"github.com/Manas8803/DigiAuth/pkg/main-app/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func CreateWallet(r *gin.Context) {

	//* Configure header for cors
	r.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	r.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	//* Check for invalid json
	var req models.Wallet
	if bindErr := r.BindJSON(&req); bindErr != nil {
		log.Println(bindErr)
		r.JSON(http.StatusBadGateway, responses.ErrorResponse{Message: "Invalid JSON"})
		return
	}

	//* Validating if all the fields are present
	if validationErr := validate.Struct(req); validationErr != nil {
		log.Println(validationErr)
		r.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "Please provide the required fields"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	queries := db.New(configs.CONN)
	//! TBD
	//* Checking whether the user is registered
	// user, userErr := queries.GetUserByEmail(ctx, req.Email)
	// if userErr != nil {
	// 	if strings.Contains(userErr.Error(), "no rows in result set") {
	// 		r.JSON(http.StatusBadGateway, responses.ErrorResponse{Message: "Email is not registered."})
	// 		return
	// 	}
	// 	log.Println(userErr)
	// 	r.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Internal server error : " + userErr.Error()})
	// 	return
	// }
	//!

	//* Generate DID
	var req_did = ledger.RegisterDIDRequest{
		Alias: "DID",
		Seed:  req.Email,
		Role:  "STEWARD",
	}
	res, registerDidErr := ledger.RegisterDID(&req_did)
	if registerDidErr != nil {
		log.Println(registerDidErr.Error())
		r.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Error in generating DID: " + registerDidErr.Error()})
		return
	}

	//* Creating Wallet
	wallet, walletErr := queries.CreateWallet(ctx, db.CreateWalletParams{
		Did:   res.DID,
		Email: req.Email,
	})
	if walletErr != nil {
		log.Println(walletErr)
		r.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Internal server error : " + walletErr.Error()})
		return
	}

	r.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully created Wallet", Data: map[string]interface{}{"Wallet-ID": wallet.ID}})
}

func IssueCeritificate(r *gin.Context) {

}
