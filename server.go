package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"./app/routes"
)

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.ListenAndServe("0.0.0.0:8080", routes.Include())

}
