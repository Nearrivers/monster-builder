package main

import (
	"fmt"
	"html/template"
	"log"
	"nearrivers/monster-creator/src/campaign"
	"nearrivers/monster-creator/src/db"
	"nearrivers/monster-creator/src/monster"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setupRoutes(r *httprouter.Router) {
	r.ServeFiles("/styles/*filepath", http.Dir("./styles"))
	r.ServeFiles("/scripts/*filepath", http.Dir("./scripts"))

	indexRoute := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	reactivityRoute := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		defer r.Body.Close()
		r.ParseForm()

		body := r.FormValue("test")

		fmt.Fprintf(w, "%s", body)
	}

	r.GET("/", indexRoute)
	r.POST("/", reactivityRoute)

	campaign.ConfigureRoutes(r)
	monster.ConfigureRoutes(r)
}

func main() {
	router := httprouter.New()

	err := db.ConnectToDb()

	if err != nil {
		panic(err)
	}

	setupRoutes(router)

	fmt.Println("Server launched")

	// log.Fatal permet de, si jamais une erreur advient dans la fonction http.ListenAndServe(), logger cette erreur et terminer le programme
	log.Fatal(http.ListenAndServe(":8080", router))
}
