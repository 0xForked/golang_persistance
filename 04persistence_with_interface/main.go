package main


import (
	"awesomeGoProject/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.MakeDb( "root:root@(127.0.0.1:3306)/awesome_go?parseTime=true")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db}

	r := mux.NewRouter()

	r.HandleFunc("/examples", env.AllExample).Methods("GET") // basic using env.example

	log.Fatal(http.ListenAndServe(":8083", r))
}

func (env *Env)  AllExample(w http.ResponseWriter, r *http.Request) {
	tst, err := env.db.AllDataInterface()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range tst {
		_, _ = fmt.Fprintf(w, "%s %s,\n", bk.Id, bk.Title)
	}
}