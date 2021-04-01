package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"stjonathanmark.com/people/data"
	"stjonathanmark.com/people/models"

	"github.com/gorilla/mux"
)

func GetPersons(db *data.DataSource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pageNumber := getIntQueryParam(r, "pageNumber")
		pageSize := getIntQueryParam(r, "pageSize")

		if pageNumber == 0 && pageSize > 0 {
			pageSize = 0
		}

		persons, err := db.GetPersons(Offset(pageNumber, pageSize), pageSize)
		HandleError(w, err, "API - Get Persons Error ", http.StatusInternalServerError)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}
}

func GetPerson(db *data.DataSource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, parseErr := strconv.ParseInt(params["id"], 0, 64)
		HandleError(w, parseErr, "API - Invalid Parameter Data Type Error ", http.StatusBadRequest)

		person, dataErr := db.GetPerson(id)
		HandleError(w, dataErr, "API - Get Person Error ", http.StatusInternalServerError)

		if person.Id != id {
			HandleError(w, Error{"Person Not Found"}, "API - Person Not Found", http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(person)
	}
}

func CreatePerson(db *data.DataSource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var person models.Person

		decodeErr := json.NewDecoder(r.Body).Decode(&person)
		HandleError(w, decodeErr, "API - Invalid Json Format/Model Error ", http.StatusBadRequest)

		createErr := db.CreatePerson(&person)
		HandleError(w, createErr, "API - Create Person Error ", http.StatusInternalServerError)

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdatePerson(db *data.DataSource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var person models.Person

		id, parseErr := strconv.ParseInt(params["id"], 0, 64)
		HandleError(w, parseErr, "API - Invalid Parameter(id) Data Type Error ", http.StatusBadRequest)

		decodeErr := json.NewDecoder(r.Body).Decode(&person)
		HandleError(w, decodeErr, "API - Invalid Json Format/Model Error ", http.StatusBadRequest)

		person.Id = id
		updateErr := db.UpdatePerson(&person)
		HandleError(w, updateErr, "API - Update Person Error ", http.StatusInternalServerError)

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeletePerson(db *data.DataSource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, parseErr := strconv.ParseInt(params["id"], 0, 64)
		HandleError(w, parseErr, "API - Invalid Parameter(id) Data Type Error ", http.StatusBadRequest)

		deleteErr := db.DeletePerson(id)
		HandleError(w, deleteErr, "API - Delete Person Error ", http.StatusInternalServerError)

		w.WriteHeader(http.StatusNoContent)
	}
}

func AddPersonHandlers(router *mux.Router, db *data.DataSource) {
	router.HandleFunc("/api/person", GetPersons(db)).Methods("GET")
	router.HandleFunc("/api/person/{id}", GetPerson(db)).Methods("GET")
	router.HandleFunc("/api/person", CreatePerson(db)).Methods("POST")
	router.HandleFunc("/api/person/{id}", UpdatePerson(db)).Methods("PUT")
	router.HandleFunc("/api/person/{id}", DeletePerson(db)).Methods("DELETE")
}
