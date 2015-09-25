package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"

	"go-aio/bin/response"

	"go-aio/app/models"
)

type (
	FraseController struct {
		session *mgo.Session
	}
)

func NewFraseController(s *mgo.Session) *FraseController {
	return &FraseController{s}
}

func (fc FraseController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f := models.Frase{}

	json.NewDecoder(r.Body).Decode(&f)

	f.Id = bson.NewObjectId()

	fc.session.DB("db_aio").C("frases").Insert(f)

	response.Json(w, f)

}

func (fc FraseController) GetAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var lista []models.Frase

	if err := fc.session.DB("db_aio").C("frases").Find(nil).All(&lista); err != nil {
		w.WriteHeader(404)
		return
	}

	response.Json(w, lista)

}

func (fc FraseController) GetBy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	f := models.Frase{}

	if err := fc.session.DB("db_aio").C("frases").FindId(oid).One(&f); err != nil {
		w.WriteHeader(404)
		return
	}
	response.Json(w, f)

}

func (fc FraseController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := fc.session.DB("db_aio").C("frases").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)

}
