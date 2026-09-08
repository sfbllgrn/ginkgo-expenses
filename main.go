package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:   "My Home Server",
			Message: "Hello from Go!",
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
	})

	log.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
