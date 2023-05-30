package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"backend/lambda-golang-gambituser/awsgo"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

/*
Ejecución de la Lambda
*/
func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !ValidateParams() {
		fmt.Println("Error en parametros, enviar 'SecretName'")
		err := errors.New("Error en parametros, enviar 'SecretName'")
		return event, err
	}
}

/*
Validate si el SecretName del SecretManager es devuelto por AWS
*/
func ValidateParams() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}
