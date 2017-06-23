package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Task string
	Done bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("views/todo.html"))
	todos := []Todo{

		{"learn go", false},
		{"learn angular", true},
		{"read examples", false},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Todos []Todo }{todos})
	})
	http.ListenAndServe(":8080", nil)

}
