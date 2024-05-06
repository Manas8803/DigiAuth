package ledger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ldej/go-acapy-client"
)

type RegisterDIDRequest struct {
	Alias string `json:"alias" validate:"required,min=5,max=50"`
	Seed  string `json:"seed" validate:"required,min=5"`
	Role  string `json:"role" validate:"required,oneof=STEWARD ENDORSER"`
	DID   string `json:"did"`
}

func RegisterDID(req *RegisterDIDRequest) (*acapy.RegisterDIDResponse, error) {

	ledgerURL := os.Getenv("LEDGER_URL") + ":9000/"

	var res acapy.RegisterDIDResponse
	body, _ := json.Marshal(req)
	resp, _ := http.Post(ledgerURL+"register", "application/json", bytes.NewBuffer(body))
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return &acapy.RegisterDIDResponse{}, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("Error unmarshaling response body:", err)
		return &acapy.RegisterDIDResponse{}, err
	}

	return &res, nil
}

//^ HOW TO CALL THIS FUNCTION :
/*
res, err := ledger.RegisterDID("firstuser", "firstuserseed", "STEWARD")
	if err != nil {
		log.Println("Error in registering DID", err)
		return
	}
	log.Println(res)
*/
