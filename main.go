package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"namlk/gomysql/config"
	"namlk/gomysql/controler"
	"namlk/gomysql/entities"
	"namlk/gomysql/helper"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db            *sql.DB
	tmpl          *template.Template
	err           error
	todoControler controler.TodoControler
)

func writeMessage(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
}

func handleGreeting(w http.ResponseWriter, r *http.Request) {
	writeMessage(w, "Hello world. And I wish all thing be fine for my family, my friends and everyone in the world.")
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	todos := todoControler.GetAllTodos()
	err = tmpl.ExecuteTemplate(w, "home.html", todos)
	helper.ErrCheck(err)
}

func handleRenderCreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "createForm.html", nil)
}

func handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	helper.ErrCheck(err)
	todoName := r.FormValue("todoName")
	expired := r.FormValue("todoExpired")
	completed := r.FormValue("todoCompleted")
	var isExpired, isCompleted bool
	if expired != "" {
		isExpired = true
	} else {
		isExpired = false
	}
	if completed != "" {
		isCompleted = true
	} else {
		isCompleted = false
	}
	err = todoControler.CreateNewTodo(todoName, isExpired, isCompleted)
	if err != nil {
		tmpl.ExecuteTemplate(w, "createForm.html", "Have an error when you try to add new tod. Try again!")
	} else {
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}

func handleRenderUpdateForm(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	helper.ErrCheck(err)
	todo, err := todoControler.GetTodoById(id)
	response := make(map[string]interface{})
	if err != nil {
		err = tmpl.ExecuteTemplate(w, "notFound.html", id)
		helper.ErrCheck(err)
	} else {
		response["todo"] = todo
		err = tmpl.ExecuteTemplate(w, "updateForm.html", response)
		helper.ErrCheck(err)
	}
}

func handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	err = r.ParseForm()
	helper.ErrCheck(err)
	idStr := r.FormValue("todoId")
	id, err := strconv.Atoi(idStr)
	helper.ErrCheck(err)
	name := r.FormValue("todoName")
	expired := r.FormValue("todoExpired")
	completed := r.FormValue("todoCompleted")
	var isExpired, isCompleted bool
	if expired != "" {
		isExpired = true
	} else {
		isExpired = false
	}
	if completed != "" {
		isCompleted = true
	} else {
		isCompleted = false
	}
	err = todoControler.UpdateTodo(entities.NewTodo(id, name, isExpired, isCompleted))
	response := make(map[string]interface{})
	if err != nil {
		response["message"] = "Some fields can't valid. Try again."
		err = tmpl.ExecuteTemplate(w, "updateForm.html", response)
		helper.ErrCheck(err)
	} else {
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}

func handleRenderConfirmDeleteTodo(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	helper.ErrCheck(err)
	todo, err := todoControler.GetTodoById(id)
	if err != nil {
		err = tmpl.ExecuteTemplate(w, "notFound.html", id)
		helper.ErrCheck(err)
	} else {
		err = tmpl.ExecuteTemplate(w, "confirmDelete.html", todo)
		helper.ErrCheck(err)
	}
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	helper.ErrCheck(err)
	err = todoControler.DeleteTodo(id)
	helper.ErrCheck(err)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func main() {
	// Contructor db connection
	db, err = config.GetMySQLDB()
	helper.ErrCheck(err)

	// Close database connection
	defer db.Close()

	// Contructor controler for todo
	todoControler = controler.TodoControler{DB: db}

	// Parse templates html
	tmpl, err = template.ParseGlob("templates/*.html")
	helper.ErrCheck(err)

	//Handle function and router
	http.HandleFunc("/", handleGreeting)
	http.HandleFunc("/home", handleHome)
	http.HandleFunc("/form-create", handleRenderCreateForm)
	http.HandleFunc("/create", handleCreateTodo)
	http.HandleFunc("/form-update/", handleRenderUpdateForm)
	http.HandleFunc("/update", handleUpdateTodo)
	http.HandleFunc("/delete/", handleRenderConfirmDeleteTodo)
	http.HandleFunc("/delete-confirm/", handleDeleteTodo)

	// Start and
	fmt.Println("Server has starting and listening in port 80...")
	err = http.ListenAndServe(":80", nil)
	helper.ErrCheck(err)
}
