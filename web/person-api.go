package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"stjonathanmark.com/people/data"
	"stjonathanmark.com/people/models"

	"github.com/gorilla/mux"
)

var db = data.NewDataSource()

func GetPersons(w http.ResponseWriter, r *http.Request) {
	pageNumber := getIntQueryParam(r, "pageNumber")
	pageSize := getIntQueryParam(r, "pageSize")

	if pageNumber == 0 && pageSize > 0 {
		pageSize = 0
	}

	persons, err := db.GetPersons(Offset(pageNumber, pageSize), pageSize)
	if HandleError(w, err, "API - Get Persons Error - ", http.StatusInternalServerError) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, parseErr := strconv.ParseInt(params["id"], 0, 64)
	if HandleError(w, parseErr, "API - Invalid Parameter Data Type Error - ", http.StatusBadRequest) {
		return
	}

	person, dataErr := db.GetPerson(id)

	if dataErr != nil && person.Id != id {
		HandleError(w, Error{fmt.Sprintf("Person record with id \"%v\" does not exist in database.", id)}, "API - Person Not Found - ", http.StatusNotFound)
		return
	} else if dataErr != nil {
		HandleError(w, dataErr, "API - Get Person Error - ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	decodeErr := json.NewDecoder(r.Body).Decode(&person)
	if HandleError(w, decodeErr, "API - Invalid Json Format/Model Error - ", http.StatusBadRequest) {
		return
	}

	createErr := db.CreatePerson(&person)
	if HandleError(w, createErr, "API - Create Person Error - ", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person

	id, parseErr := strconv.ParseInt(params["id"], 0, 64)
	if HandleError(w, parseErr, "API - Invalid Parameter(id) Data Type Error - ", http.StatusBadRequest) {
		return
	}

	decodeErr := json.NewDecoder(r.Body).Decode(&person)
	if HandleError(w, decodeErr, "API - Invalid Json Format/Model Error - ", http.StatusBadRequest) {
		return
	}

	person.Id = id
	updateErr := db.UpdatePerson(&person)
	if HandleError(w, updateErr, "API - Update Person Error - ", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, parseErr := strconv.ParseInt(params["id"], 0, 64)
	if HandleError(w, parseErr, "API - Invalid Parameter(id) Data Type Error - ", http.StatusBadRequest) {
		return
	}

	deleteErr := db.DeletePerson(id)
	if HandleError(w, deleteErr, "API - Delete Person Error - ", http.StatusInternalServerError) {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AddPersonHandlers(router *mux.Router) {
	router.HandleFunc("/api/person", GetPersons).Methods("GET")
	router.HandleFunc("/api/person/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/api/person", CreatePerson).Methods("POST")
	router.HandleFunc("/api/person/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/api/person/{id}", DeletePerson).Methods("DELETE")
}
