package main

import (
	"fmt"
	"html/template"
	"log"
	"nearrivers/monster-creator/src/db"
	router "nearrivers/monster-creator/src/routers"
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

	r.GET("/", indexRoute)

	router.ConfigureCampaignRoutes(r)
}

func main() {
	router := httprouter.New()
	fmt.Println("Hello world")

	err := db.ConnectToDb()

	if err != nil {
		panic(err)
	}

	setupRoutes(router)

	// log.Fatal permet de, si jamais une erreur advient dans la fonction http.ListenAndServe(), log cette erreur et termine le programme
	log.Fatal(http.ListenAndServe(":8080", router))
}
