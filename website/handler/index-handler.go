package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("assets/index.html")

	if err != nil {
		log.Println("Error: ", err.Error())
		fmt.Fprintf(w, "Error: ", err.Error())
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println("Error: ", err.Error())
	}
}
