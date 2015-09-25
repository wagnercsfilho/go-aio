package routes

import (
	"net/http"

	"../controllers"
)

func Include() {

	http.HandleFunc("/", controllers.Show)
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/teste", controllers.Teste)

}
