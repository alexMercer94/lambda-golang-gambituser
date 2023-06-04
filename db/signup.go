package db

import (
	"backend/lambda-golang-gambituser/models"
	"backend/lambda-golang-gambituser/tools"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
Insertar datos en la tabla Users de la BD
*/
func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sqlQuery := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES('" + sig.UserEmail + "', '" + sig.UserUUID + "', '" + tools.DateMySQL() + "')"
	fmt.Println(sqlQuery)
	_, err = Db.Exec(sqlQuery)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución de inserción exitosa")

	return nil
}
