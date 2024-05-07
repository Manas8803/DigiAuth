# Removed GenerateDID func from controllers

func GenerateDID(r \*gin.Context) {

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
