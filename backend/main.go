package main

import (
	"net/http"

	"github.com/SuperTikuwa/book_app/handler"
	"github.com/SuperTikuwa/book_app/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r = r.PathPrefix("/api").Subrouter()

	r.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy."))
	})

	r.HandleFunc("/test", test)

	r.HandleFunc("/user", handler.StoreUser).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	u := model.User{
		Name:        "test",
		Email:       "test",
		CognitoUUID: gofakeit.UUID(),
	}
	if err := u.Create(); err != nil {
		w.Write([]byte(err.Error()))
		return

	}

	w.Write([]byte("OK"))
}
