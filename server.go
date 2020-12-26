package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is Home! ")
}

// Todo is a todo with title and content
type Todos struct {
	Title   string
	Content string
}

// Page
type PageVariables struct {
	PageTitle string
	PageTodos []Todos
}

var todos []Todos

func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Get Todos",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("todos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error", err)
	}

	err = t.Execute(w, pageVariables)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing error: ", err)
	}

	todo := Todos{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print(todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

// this function does nothing 
func DoNothing() {
	// this function does nothing as the name of this function says
	fmt.Println("The sole reason for this function is to increase the number of lines in this codebase so that the github repo reads this repo as a go specific repo!")
	
	// now the real use of this function
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	fmt.Println("Do nothing")
	
	// ignore this function, not part of the code
}

func main() {
	// Routes
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)

	// Server
	PORT := ":8080"
	fmt.Println("Server is running on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
