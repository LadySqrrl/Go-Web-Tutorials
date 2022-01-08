package main

import (
	"html/template"
	"net/http"
)

// Todo is a stuct that holds to do information
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData is a struct that holds the list of all Todo items
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)

}
