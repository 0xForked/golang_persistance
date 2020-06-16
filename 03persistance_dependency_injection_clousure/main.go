package main

import (
	"awesomeGoProject/models"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := models.NewDB( "root:root@(127.0.0.1:3306)/awesome_go?parseTime=true")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	r := mux.NewRouter()

	r.Handle("/examples", AllExample(env)).Methods("GET") // closure AllExample(env)

	log.Fatal(http.ListenAndServe(":8082", r))
}

func  AllExample (env *Env) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tst, err := models.AllDataDI(env.db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		for _, bk := range tst {
			_, _ = fmt.Fprintf(w, "%s %s,\n", bk.Id, bk.Title)
		}
	})
}