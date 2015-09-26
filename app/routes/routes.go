package routes

import (
	"net/http"
	. "go-aio/app/controllers"
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
)

func Include() *httprouter.Router {
	r := httprouter.New()

  r.ServeFiles("/public/*filepath", http.Dir("../go-aio/public/"))
	fraseController := NewFraseController(getSession())
	homeController := NewHomeController()

	r.GET("/", homeController.Index)

	r.GET("/frase/:id", fraseController.GetBy)
	r.GET("/frase", fraseController.GetAll)
	r.POST("/frase", fraseController.Create)
	r.DELETE("/frase/:id", fraseController.Delete)

	return r
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
