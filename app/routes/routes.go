package routes

import (
    "net/http"
    "../controllers"
)

func Include() {
    
    http.HandleFunc("/", controllers.Show);
       
}