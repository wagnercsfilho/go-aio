package main

import (
	"net/http"

	"./app/routes"
)

func main() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.ListenAndServe("0.0.0.0:8080", routes.Include())

}
