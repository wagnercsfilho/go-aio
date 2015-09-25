package response

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func Json(w http.ResponseWriter, v interface{}) {

	data, err := json.MarshalIndent(v, "  ", " ")
	CheckError(err, "Falha no parser json")

	fmt.Fprintf(w, string(data))

}

func View(w http.ResponseWriter, file string, v interface{}) {

	root_path, err := os.Getwd()

	t, err := template.ParseFiles(root_path + "/app/views/" + file + ".html")
	CheckError(err, "Falha no parser html")

	t.Execute(w, nil)

}

func CheckError(err error, msn string) {
	if err != nil {
		log.Fatalln(err, msn)
	}
}
