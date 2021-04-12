package main

import (
	"log"
	"net/http"
	"os"

	"stjonathanmark.com/people/web"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	web.AddPersonHandlers(router)

	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, router)

	log.Fatal("Error occurred starting web sever ", err)
}
