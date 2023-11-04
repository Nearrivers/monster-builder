package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	styleFs := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles", styleFs))

	scriptsFs := http.FileServer(http.Dir("./scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts", scriptsFs))

	http.HandleFunc("/", h1)

	// log.Fatal permet de, si jamais une erreur advient dans la fonction http.ListenAndServe(), log cette erreur et termine le programme
	log.Fatal(http.ListenAndServe(":8080", nil))
}
