package main

import (
	"net/http"

	"./app/routes"
)

func main() {
	http.Handle("/*", http.FileServer(http.Dir("./public/*.*")))
	http.ListenAndServe("0.0.0.0:8080", routes.Include())

}
