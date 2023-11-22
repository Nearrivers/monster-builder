package campaign

import (
	"fmt"
	"html/template"
	"nearrivers/monster-creator/src/models"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type templateData struct {
	Campaigns []models.Campaign
}

const fileBasePath = "./templates/campaign/"

func GetAllCampaignsTemplate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tD templateData
	var campaigns []models.Campaign
	campaigns, err := GetAllCampaigns()

	if err != nil {
		http.Error(w, "Aucune campagne trouvée", http.StatusNotFound)
	}

	tD.Campaigns = campaigns
	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "AllCampaigns.html")))
	t.Execute(w, tD)
}

func GetAllCampaignsSelect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles("./templates/monster/NewMonster.html"))
	var tD templateData
	campaigns, err := GetAllCampaigns()

	if err != nil {
		http.Error(w, "Aucune campagne trouvée", http.StatusNotFound)
	}

	tD.Campaigns = campaigns
	t.ExecuteTemplate(w, "campaign-select", tD)
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
