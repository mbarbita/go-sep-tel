package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/context"

	cfgutils "github.com/mbarbita/golib-cfgutils"
)

func home(w http.ResponseWriter, r *http.Request) {

	tData := r.Host

	// Execute template
	err := htmlTmpl.ExecuteTemplate(w, "index.html", tData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

var cfgMap = cfgutils.ReadCfgFile("cfg.ini", false)

var telMap = cfgutils.ReadCfgFile(cfgMap["tel"], false)
var htmlTmpl = template.Must(template.ParseGlob("templates/*.*"))

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/msg", wsMessage)

	log.Println("Running...")

	err := http.ListenAndServe(cfgMap["server"],
		context.ClearHandler(http.DefaultServeMux))

	if err != nil {
		panic("ListenAndServe err: " + err.Error())
	}
}
