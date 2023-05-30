package db

import (
	awsecretm "backend/lambda-golang-gambituser/awssecretm"
	"backend/lambda-golang-gambituser/models"
	"os"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = awsecretm.GetSecret(os.Getenv("SecretName"))
	return err
}
