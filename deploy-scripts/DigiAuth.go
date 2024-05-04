package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/joho/godotenv"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DigiAuthStackProps struct {
	awscdk.StackProps
}

func CreateDigiAuthStack(scope constructs.Construct, id string, props *DigiAuthStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	auth_handler := awslambda.NewFunction(stack, jsii.String("Auth"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("")),
		Runtime: awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
		},
	})

	awsapigateway.NewLambdaRestApi(stack, jsii.String("authTest"), &awsapigateway.LambdaRestApiProps{
		Handler: auth_handler,
	})


	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	CreateDigiAuthStack(app, "DeployScriptsStack", &DigiAuthStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalln("Error loading .env file : ", err)
	}

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
