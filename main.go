package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"backend/lambda-golang-gambituser/awsgo"
	"backend/lambda-golang-gambituser/db"
	"backend/lambda-golang-gambituser/models"

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

	var data models.SignUp
	/*
		Obtener data del evento de Cognito y setearla en el modelo
	*/
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)

		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}

	}

	// Leer el Secret de Secret Manager
	err := db.ReadSecret()
	if err != nil {
		fmt.Println(" > Error al leer el Secret de AWS Secret Manager: " + err.Error())
		return event, err
	}

	// Insertar datos en la Tabla User de la BD
	err = db.SignUp(data)
	return event, err
}

/*
Validate si el SecretName del SecretManager es devuelto por AWS
*/
func ValidateParams() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}
