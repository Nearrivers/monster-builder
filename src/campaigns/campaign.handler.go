package campaign

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetAllCampaigns(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Toutes les campagnes \n")
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
