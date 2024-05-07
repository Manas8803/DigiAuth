package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DigiAuthStackProps struct {
	awscdk.StackProps
}

const APP_NAME="DigiAuth"
func CreateDigiAuthStack(scope constructs.Construct, id string, props *DigiAuthStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	wallet_handler := awslambda.NewFunction(stack, jsii.String("Wallet"), &awslambda.FunctionProps{
		FunctionName: jsii.String(APP_NAME+"Wallet"),
		Code:    awslambda.Code_FromAsset(jsii.String("../app"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(20)),
		Environment: &map[string]*string{
			"RELEASE_MODE": jsii.String("prod"),
			"LEDGER_URL": jsii.String(os.Getenv("LEDGER_URL")),
		},
	})

	awsapigateway.NewLambdaRestApi(stack, jsii.String("Wallet_Gateway"), &awsapigateway.LambdaRestApiProps{
		Handler: wallet_handler,
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
		},
	})


	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	CreateDigiAuthStack(app, "DigiAuthStack", &DigiAuthStackProps{
		awscdk.StackProps{
			StackName: jsii.String("DigiAuthStack"),
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatalln("Error loading .env file : ", err)
	// }

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
