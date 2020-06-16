package main

import (
	"awesomeGoProject/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDB struct{}

func (mdb *mockDB) AllDataInterface() ([]*models.Test, error) {
	tst := make([]*models.Test, 0)
	tst = append(tst, &models.Test{Id: "11", Title: "Emma"})
	tst = append(tst, &models.Test{Id: "22", Title: "The Time Machine"})
	return tst, nil
}

func TestDataIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/data", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.AllExample).ServeHTTP(rec, req)

	expected := "11 Emma,\n22 The Time Machine,\n"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}