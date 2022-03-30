package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Schema struct {
	Variable string
}

func handleScripts(w http.ResponseWriter, r *http.Request) {

	k6Template := readTemplate()
	tmpl := template.Must(template.New("K6Script").Parse(k6Template))
	tmpl.Execute(w, Schema{Variable: "Kubernetes"})

}

func main() {

	// define routes
	http.HandleFunc("/scripts", handleScripts)

	// starting
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func readTemplate() string {
	data, err := ioutil.ReadFile("k6.js")
	if err != nil {
		fmt.Printf("File reading error")
	}
	return string(data)
}
