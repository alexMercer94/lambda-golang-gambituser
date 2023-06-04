package tools

import (
	"fmt"
	"time"
)

/*
Formatear fecha para guardar en MySQL
*/
func DateMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%2d:%2d:%2d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}