package main

import (
	"net/http"

	"github.com/carlhueffmeier/gotodo/todo"
)

func main() {
	todoService := todo.NewService()
	todoAPI := todo.NewAPI(&todoService)
	http.Handle("/todos/", http.StripPrefix("/todos/", todoAPI))
	http.Handle("/todos", http.StripPrefix("/todos", todoAPI))
	http.ListenAndServe(":8880", nil)
}
