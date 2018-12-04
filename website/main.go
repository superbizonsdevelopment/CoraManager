package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/superbizonsdevelopment/CoraManager/api"

	"./manager"

	"./handler"
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

	log.Println("Loading Configuration file!")

	_, err = manager.LoadConfiguration()

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	log.Println("Successfully loaded configuration file!")

	u := api.User{"Antoni Bednarski", "XDD", "antoni.p.bednarski@gmail.com", "admin", nil}
	err = manager.SendMail(&u)

	if err != nil {
		log.Println("Error", err.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleIndex)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":5000", mux)
}
