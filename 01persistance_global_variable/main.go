package main

import (
	"awesomeGoProject/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	models.InitDB( "root:root@(127.0.0.1:3306)/awesome_go?parseTime=true")

	r := mux.NewRouter()

	r.HandleFunc("/examples", AllExample).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func AllExample(w http.ResponseWriter, r *http.Request) {
	tst, err := models.AllDataGV()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range tst {
		_, _ = fmt.Fprintf(w, "%s %s,\n", bk.Id, bk.Title)
	}
}
