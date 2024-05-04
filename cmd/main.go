package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/basic/docs"
)

var ginLambda *ginadapter.GinLambda

func init() {

	prod := os.Getenv("RELEASE_MODE")
	if prod == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	api := router.Group("/api/v1")
	
	routes.LedgerRoute(api)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
