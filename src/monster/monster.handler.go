package monster

import (
	"fmt"
	"html/template"
	"nearrivers/monster-creator/src/db"
	"nearrivers/monster-creator/src/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
)

const fileBasePath = "./templates/monster/"

type monsterTemplateData struct {
	Monsters []models.Monster
}

type traitsDto struct {
	Name        string
	Description string
}

type actionDto struct {
	Name        string
	Description string
	Type        string
}

type monsterDto struct {
	Campaign          uint8
	Name              string
	Type              string
	SubType           string
	Height            string
	Alignment         string
	ArmorClass        string
	HealthPoints      uint64
	Speed             uint32
	Strength          uint8
	Dexterity         uint8
	Constitution      uint8
	Intelligence      uint8
	Wisdom            uint8
	Charisma          uint8
	SavingThrows      string
	Skills            string
	Vulnerabilities   string
	Resistances       string
	DamageImmunities  string
	StateImmunities   string
	Senses            string
	PassivePerception uint8
	Languages         string
	Challenge         string
	MasteryBonus      uint8
	SpecialTraits     []traitsDto
	Actions           []actionDto
	Reactions         []actionDto
	BonusActions      []actionDto
	LegendaryActions  []actionDto
	Description       string
	Portrait          []byte
}

var decoder = schema.NewDecoder()

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

func getNewMonsterActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "/fragments/NewMonster/Actions.html")))
	skillId, err := strconv.Atoi(ps.ByName("id"))
	actionType := ps.ByName("type")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := map[string]interface{}{
		"current": skillId,
		"next":    skillId + 1,
		"type":    actionType,
	}

	t.Execute(w, data)
}

func createMonster(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var monster monsterDto

	err = decoder.Decode(&monster, r.PostForm)
	if err != nil {
		http.Error(w, "Erreur lors du décodage : "+err.Error(), http.StatusBadRequest)
		return
	}

	var monsterEntity models.Monster
	db := db.GetDbConnection()

	monsterEntity.Name = monster.Name
	monsterEntity.Type = monster.Type
	monsterEntity.SubType = monster.SubType
	monsterEntity.Height = monster.Height
	monsterEntity.Alignment = monster.Alignment
	monsterEntity.ArmorClass = monster.ArmorClass
	monsterEntity.HealthPoints = monster.HealthPoints
	monsterEntity.Speed = monster.Speed
	monsterEntity.Strength = monster.Strength
	monsterEntity.Dexterity = monster.Dexterity
	monsterEntity.Constitution = monster.Constitution
	monsterEntity.Intelligence = monster.Intelligence
	monsterEntity.Wisdom = monster.Wisdom
	monsterEntity.Charisma = monster.Charisma
	monsterEntity.SavingThrows = monster.SavingThrows
	monsterEntity.Skills = monster.Skills
	monsterEntity.Vulnerabilities = monster.Vulnerabilities
	monsterEntity.Resistances = monster.Resistances
	monsterEntity.DamageImmunities = monster.DamageImmunities
	monsterEntity.StateImmunities = monster.StateImmunities
	monsterEntity.Senses = monster.Senses
	monsterEntity.PassivePerception = monster.PassivePerception
	monsterEntity.Languages = monster.Languages
	monsterEntity.Challenge = monster.Challenge
	monsterEntity.MasteryBonus = monster.MasteryBonus
	monsterEntity.Description = monster.Description
	monsterEntity.Portrait = monster.Portrait

	specialTraits := make([]models.SpecialTrait, len(monster.SpecialTraits))
	if len(monster.SpecialTraits) > 0 {
		for _, st := range monster.SpecialTraits {
			newSpecialTrait := models.SpecialTrait{}
			newSpecialTrait.Name = st.Name
			newSpecialTrait.Description = st.Description
		}
	}

	monsterEntity.SpecialTraits = specialTraits

	actions := make([]models.Action, len(monster.Actions))
	reactions := make([]models.Action, len(monster.Reactions))
	bonusactions := make([]models.Action, len(monster.BonusActions))
	legendaryActions := make([]models.Action, len(monster.LegendaryActions))

	if len(monster.Actions) > 0 {
		for _, r := range monster.Reactions {
			newAction := models.Action{}
			newAction.Name = r.Name
			newAction.Description = r.Description
			newAction.Type = "Action"
			actions = append(actions, newAction)
		}
	}

	if len(monster.Reactions) > 0 {
		for _, r := range monster.Reactions {
			newReaction := models.Action{}
			newReaction.Name = r.Name
			newReaction.Description = r.Description
			newReaction.Type = "Réaction"
			reactions = append(reactions, newReaction)
		}
	}

	if len(monster.BonusActions) > 0 {
		for _, r := range monster.BonusActions {
			newBonusAction := models.Action{}
			newBonusAction.Name = r.Name
			newBonusAction.Description = r.Description
			newBonusAction.Type = "Action bonus"
			bonusactions = append(bonusactions, newBonusAction)
		}
	}

	if len(monster.LegendaryActions) > 0 {
		for _, r := range monster.LegendaryActions {
			newLegendaryAction := models.Action{}
			newLegendaryAction.Name = r.Name
			newLegendaryAction.Description = r.Description
			newLegendaryAction.Type = "Action légendaires"
			legendaryActions = append(legendaryActions, newLegendaryAction)
		}
	}

	monsterEntity.Actions = append(actions, reactions...)
	monsterEntity.Actions = append(monsterEntity.Actions, bonusactions...)
	monsterEntity.Actions = append(monsterEntity.Actions, legendaryActions...)

	if result := db.Create(&monsterEntity); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Monstre créé")
}

func deleteCurrentAbility(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "")
}
