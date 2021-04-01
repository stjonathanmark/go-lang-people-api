package main

import (
	"log"
	"net/http"
	"os"

	"stjonathanmark.com/people/data"
	"stjonathanmark.com/people/web"

	"github.com/gorilla/mux"
)

func main() {
	connStr := os.Getenv("MSSQL_CONN_STRING")
	db := data.NewDataSource(connStr)

	router := mux.NewRouter()

	web.AddPersonHandlers(router, db)

	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, router)

	log.Fatal("Error occurred starting web sever ", err)
}
