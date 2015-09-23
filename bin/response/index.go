package response

import (
    "encoding/json"
    "fmt"
    "net/http"
    "html/template"
    "os"
)



func Json(w http.ResponseWriter, v interface{}) {
    
    data, err := json.MarshalIndent(v, "", "")
    if err != nil{
        panic(err)
    }
    
    fmt.Fprintf(w, string(data))

}

func View(w http.ResponseWriter, file string, v interface{}) {
    
    root_path, err := os.Getwd()
    
    t, err := template.ParseFiles(root_path + "/app/views/" + file + ".html")
    if err != nil {
        panic(err)
    }
    
    t.Execute(w, nil)
    
}