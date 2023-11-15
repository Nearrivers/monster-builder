package campaign

import (
	"fmt"
	"html/template"
	"nearrivers/monster-creator/src/db"
	"nearrivers/monster-creator/src/models"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type templateData struct {
	Campaigns []models.Campaign
}

const fileBasePath = "./templates/campaign/"

func GetAllCampaigns(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tD templateData
	var campaigns []models.Campaign
	db := db.GetDbConnection()

	if result := db.Find(&campaigns); result.Error != nil {
		fmt.Println(result.Error)
	}

	if length := len(campaigns); length == 0 {
		htmlStr := "<p>Aucune campagne n'a été trouvée</p>"
		tmpl, err := template.New("Not found").Parse(htmlStr)

		if err != nil {
			panic(err)
		}

		tmpl.Execute(w, nil)
	} else {
		tD.Campaigns = campaigns
		t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "AllCampaigns.html")))
		t.Execute(w, tD)
	}
}

func GetOneCampaign(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Une campagne, %s!\n", ps.ByName("id"))
}

func CreateCampaign(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	fmt.Print(r.Body, "Nouvelle campagne \n")
}

func UpdateCampaign(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()
	fmt.Print(r.Body, "Mise à jour de la campagne \n")
}

func DeleteOneCampaign(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Campagne supprimée, %s!\n", ps.ByName("id"))
}
