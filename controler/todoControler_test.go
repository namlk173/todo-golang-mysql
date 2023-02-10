package controler_test

import (
	"database/sql"
	"fmt"
	"namlk/gomysql/config"
	"namlk/gomysql/controler"
	"namlk/gomysql/helper"
	"testing"
)

var db *sql.DB
var todoControler controler.TodoControler
var err error

func TestGetTodoById(t *testing.T) {
	db, err = config.GetMySQLDB()
	helper.ErrCheck(err)
	todoControler = controler.TodoControler{DB: db}
	todo, err := todoControler.GetTodoById(1)
	fmt.Println(todo, err)
	if err != nil {
		t.Error("Fail")
	}
}
