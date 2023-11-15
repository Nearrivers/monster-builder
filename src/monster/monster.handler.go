package monster

import (
	"fmt"
	"html/template"
	"nearrivers/monster-creator/src/db"
	"nearrivers/monster-creator/src/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type monsterTemplateData struct {
	Monsters []models.Monster
}

const fileBasePath = "./templates/monster/"

func executeAbilitiesTemplates(t *template.Template, w http.ResponseWriter, ps httprouter.Params) {
	skillId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := map[string]interface{}{
		"current": skillId,
		"next":    skillId + 1,
	}

	t.Execute(w, data)
}

func getAllMonsters(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tD monsterTemplateData
	var monsters []models.Monster
	db := db.GetDbConnection()

	if result := db.Find(&monsters); result.Error != nil {
		fmt.Println(result.Error)
	}

	if length := len(monsters); length == 0 {
		htmlStr := "<p>Aucun monstre n'a été trouvé</p>"
		tmpl, err := template.New("Not found").Parse(htmlStr)

		if err != nil {
			panic(err)
		}

		tmpl.Execute(w, nil)
	} else {
		tD.Monsters = monsters
		t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "AllMonsters.html")))
		t.Execute(w, tD)
	}
}

func getNewMonsterTemplate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "NewMonster.html")))
	t.Execute(w, nil)
}

func getNewMonsterStats(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "/fragments/NewMonster/Stats.html")))
	t.Execute(w, nil)
}

func getNewMonsterSkills(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "/fragments/NewMonster/Skills.html")))
	executeAbilitiesTemplates(t, w, ps)
}

func getNewMonsterActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "/fragments/NewMonster/Actions.html")))
	executeAbilitiesTemplates(t, w, ps)
}

func deleteCurrentAbility(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "")
}
