package main

import (
	"fmt"
	"html/template"
	"log"
	campaign "nearrivers/monster-creator/src/campaigns"
	"nearrivers/monster-creator/src/db"
	"nearrivers/monster-creator/src/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setupRoutes(r *httprouter.Router) {
	styleFs := http.FileServer(http.Dir("./styles"))
	http.Handle("/styles/", http.StripPrefix("/styles", styleFs))

	scriptsFs := http.FileServer(http.Dir("./scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts", scriptsFs))

	indexRoute := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	r.GET("/", indexRoute)

	campaign.ConfigureCampaignRoutes(r)
}

func main() {
	router := httprouter.New()
	fmt.Println("Hello world")

	db, err := db.ConnectToDb()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Campaign{}, &models.Monster{}, &models.SpecialTrait{}, &models.Action{})

	setupRoutes(router)

	// log.Fatal permet de, si jamais une erreur advient dans la fonction http.ListenAndServe(), log cette erreur et termine le programme
	log.Fatal(http.ListenAndServe(":8080", router))
}
