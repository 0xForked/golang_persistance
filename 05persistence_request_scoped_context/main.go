package main

import (
	"awesomeGoProject/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

func main() {
	db, err := models.NewDB("root:root@(127.0.0.1:3306)/awesome_go?parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	ctx := context.WithValue(context.Background(), "db", db)

	r := mux.NewRouter()

	r.Handle("/examples", &ContextInjector{ctx, http.HandlerFunc(AllExample)}).Methods("GET") // basic using env.example

	log.Fatal(http.ListenAndServe(":8084", r))
}

func AllExample(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	tst, err := models.AllDataDI(db)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range tst {
		_, _ = fmt.Fprintf(w, "%s %s,\n", bk.Id, bk.Title)
	}
}
