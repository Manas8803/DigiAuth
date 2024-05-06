package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Manas8803/DigiAuth/pkg/lib/ledger"
	"github.com/Manas8803/DigiAuth/pkg/main-app/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func CreateWallet(r *gin.Context) {

	//! MAKE CHANGES HERE

	r.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully created Wallet", Data: map[string]interface{}{"Some Data": "Data1"}}) //! CHANGE THIS
}
func GenerateDID(r *gin.Context) {

	var req ledger.RegisterDIDRequest
	if bindErr := r.BindJSON(&req); bindErr != nil {
		log.Println(bindErr)
		r.JSON(http.StatusBadGateway, responses.ErrorResponse{Message: "Invalid JSON"})
		return
	}

	if validationErr := validate.Struct(req); validationErr != nil {
		log.Println(validationErr)
		if strings.Contains("Key: 'RegisterDIDRequest.Role' Error:Field validation for 'Role' failed on the 'oneof' tag", validationErr.Error()) {
			r.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Invalid Role"})
			return
		}
		r.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: "Please provide the required fields"})
		return
	}

	res, registerDidErr := ledger.RegisterDID(&req)
	if registerDidErr != nil {
		log.Println(registerDidErr.Error())
		r.JSON(http.StatusInternalServerError, responses.ErrorResponse{Message: "Error in generating DID: " + registerDidErr.Error()})
		return
	}

	r.JSON(http.StatusOK, responses.SuccessResponse{Message: "Successfully Generated DID", Data: res})
}
