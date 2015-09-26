package main

import (
	"net/http"

	"go-aio/app/routes"
)


func main() {
	
	http.ListenAndServe("0.0.0.0:8080", routes.Include())

}
