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

func NewFraseController() *FraseController {
	return &FraseController{}
}

func (fc FraseController) CreateFrase(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	f := models.Frase{}

	json.NewDecoder(r.Body).Decode(&f)

	f.Id = bson.NewObjectId()

	fc.session.DB("db_aio").C("frases").Insert(f)

	response.Json(w, f)

}

func (fc FraseController) GetFrase(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

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
	fj, _ := json.Marshal(f)

	response.Json(w, fj)

}
