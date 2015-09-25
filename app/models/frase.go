package models

import "gopkg.in/mgo.v2/bson"

type (
	Frase struct {
		Id        bson.ObjectId `json:"id" bson:"_id"`
		Descricao string        `json:"descricao" bson:"descricao"`
	}
)
type Frases []Frase
