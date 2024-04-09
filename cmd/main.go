package main

import (
	"DID/lib/ledger"
	"log"
)

func main() {
	res, err := ledger.RegisterDID("firstuser", "firstuserseed", "STEWARD")
	if err != nil {
		log.Println("Error in registering DID", err)
		return
	}
	log.Println(res)
}
