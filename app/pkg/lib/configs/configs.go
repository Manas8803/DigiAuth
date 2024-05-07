package configs

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var CONN *pgx.Conn

const connectMsg string = "\n---------------------------------------------------------------------------------------------\nConnected to DB\n---------------------------------------------------------------------------------------------"

func ConnectDB() *pgx.Conn {
	ctx := context.Background()
	uri := os.Getenv("SQLURI")
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		log.Println(err)
		return nil
	}
	CONN = conn

	log.Println(connectMsg)
	return conn
}
