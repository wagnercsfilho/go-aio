package controllers

import (
    "net/http"
    "../../bin/response"
     "../models"
)

func Index(w http.ResponseWriter, r *http.Request) {
    todos := models.Todos{
        models.Todo{Name: "Write presentation"},
        models.Todo{Name: "Host meetup"},
    }

    response.View(w, "fruits/index", todos)
}

func Show(w http.ResponseWriter, r *http.Request) {
    todos := models.Todos{
        models.Todo{Name: "Write presentation"},
        models.Todo{Name: "Host meetup"},
    }

    response.Json(w, todos)
}

