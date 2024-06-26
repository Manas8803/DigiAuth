package main

import (
	"context"
	"log"
	"os"

	"github.com/Manas8803/DigiAuth/pkg/lib/configs"
	"github.com/Manas8803/DigiAuth/pkg/main-app/routes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ginLambda *ginadapter.GinLambda

func init() {

	mode := os.Getenv("RELEASE_MODE")
	if mode != "prod" {
		return
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api/v1")

	configs.ConnectDB()

	routes.LedgerRoute(api)

	ginLambda = ginadapter.New(router)

}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("NOT ABLE TO FIND .env FILE..\nContinuing...")
	}
	mode := os.Getenv("RELEASE_MODE")
	if mode == "testing" {
		TestRun()
		return
	}

	lambda.Start(Handler)
}

func TestRun() {
	router := gin.Default()

	api := router.Group("/api/v1")

	configs.ConnectDB()

	routes.LedgerRoute(api)
	router.Run("localhost:8000")
}
