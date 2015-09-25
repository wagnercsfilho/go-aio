package routes

import (
	"net/http"
	. "../controllers"
	"gopkg.in/mgo.v2"

	"github.com/julienschmidt/httprouter"
)

func Include() *httprouter.Router {
	r := httprouter.New()

	fraseController := NewFraseController(getSession())
	homeController := NewHomeController()
	
	r.GET("/", homeController.Index)

	r.GET("/frase/:id", fraseController.GetBy)
	r.GET("/frase", fraseController.GetAll)
	r.POST("/frase", fraseController.Create)
	r.DELETE("/frase/:id", fraseController.Delete)

	// Serve static files from the ./public directory
	r.ServeFiles("/go-aio/public/*filepath", http.Dir("/public"))
	
	return r
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
