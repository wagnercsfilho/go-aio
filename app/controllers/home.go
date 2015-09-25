package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"go-aio/bin/response"
)

type (
	HomeController struct{}
)

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (fc HomeController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response.View(w, "home/index", nil)
}
