package main

import (
	"log"
	"net/http"
	"os"

	"./handler"
	"./manager"
)

func main() {

	log.Println("Starting Cora Manager Website")

	_, err := os.Stat("website.db")

	if os.IsNotExist(err) {
		_, err = os.Create("website.db")
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		log.Println("website.db file successfully created!")
	}

	_, err = os.Stat("config.json")

	if os.IsNotExist(err) {
		_, err = os.Create("config.json")
		if err != nil {
			log.Println("Error: ", err.Error())
		}
		log.Println("config.json file successfully created!")
	}

	log.Println("Connecting to database.")
	err = manager.ConnectToDatabase()

	if err != nil {
		log.Println("Error: ", err.Error())
	}

	log.Println("Successfully connected with database!")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleIndex)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":5000", mux)
}
