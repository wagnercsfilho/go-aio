package routes

import (
	"../controllers"

	"github.com/julienschmidt/httprouter"
)

func Include() *httprouter.Router {
	r := httprouter.New()

	r.GET("/frase/:id", controllers.GetFrase)

	r.POST("/frase", controllers.CreateFrase)

	return r
}
