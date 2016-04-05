package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var Verbose bool = false

func Index(w http.ResponseWriter, req *http.Request) {
	if Verbose == true {
		log.Printf("Serving index page...")
	}
	t := template.New("index.html")
	t, _ = t.ParseFiles("templates/index.html")
	files, _ := ioutil.ReadDir(".")
	err := t.Execute(w, files)
	if err != nil {
		log.Fatal(err)
	}
}

func ServeFile(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	if Verbose == true {
		log.Printf("Serving file %s...\n", vars["filename"])
	}
	http.ServeFile(w, req, vars["filename"])
}

func main() {
	address := flag.String("address", "127.0.0.1", "server address")
	port := flag.Int("port", 3000, "server port")
	verbose := flag.Bool("verbose", false, "verbose?")
	flag.Parse()

	Verbose = *verbose

	hosturi := fmt.Sprintf("%s:%d", *address, *port)

	r := mux.NewRouter()

	r.HandleFunc("/", Index)
	r.HandleFunc("/files/{filename}", ServeFile)

	http.Handle("/", r)

	log.Printf("Listening on... %s\n", hosturi)
	log.Fatal(http.ListenAndServe(hosturi, nil))
}
