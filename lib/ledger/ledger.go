package ledger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ldej/go-acapy-client"
)

type registerDIDRequest struct {
	Alias string `json:"alias"`
	Seed  string `json:"seed"`
	Role  string `json:"role"`
	DID   string `json:"did"`
}

func RegisterDID(alias string, seed string, role string) (*acapy.RegisterDIDResponse, error) {

	var request registerDIDRequest
	ledgerURL := "http://127.0.0.1:9000/"
	request = registerDIDRequest{
		Alias: alias,
		Seed:  seed, // Should be random in develop mode
		Role:  string(role),
	}

	var res acapy.RegisterDIDResponse
	body, _ := json.Marshal(request)
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
