package monster

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"nearrivers/monster-creator/src/campaign"
	"nearrivers/monster-creator/src/db"
	"nearrivers/monster-creator/src/models"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
)

const fileBasePath = "./templates/monster/"

type traitsDto struct {
	Name        string
	Description string
}

type actionDto struct {
	Name        string
	Description string
	Type        string
}

type createMonsterDto struct {
	Campaign          uint
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

type serializedMonsterDto struct {
	ID        uint
	Name      string
	Type      string
	SubType   string
	Challenge string
	Portrait  string
}

var decoder = schema.NewDecoder()

func getAllMonsters(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var monsters []models.Monster
	db := db.GetDbConnection()

	if result := db.Find(&monsters); result.Error != nil {
		fmt.Println(result.Error)
	}

	if len(monsters) == 0 {
		http.Error(w, "Aucun monstre trouvé", http.StatusNotFound)
		return
	}

	var monsterDtos []serializedMonsterDto
	for _, monster := range monsters {
		var b64Portrait string

		mimeType := http.DetectContentType(monster.Portrait)

		switch mimeType {
		case "image/jpeg":
			b64Portrait += "data:image/jpeg;base64,"
		case "image/png":
			b64Portrait += "data:image/png;base64,"
		case "image/webp":
			b64Portrait += "data:image/webp;base64,"
		}

		b64Portrait += base64.StdEncoding.EncodeToString(monster.Portrait)

		dto := serializedMonsterDto{
			ID:        monster.ID,
			Name:      monster.Name,
			Type:      monster.Type,
			SubType:   monster.SubType,
			Challenge: monster.Challenge,
			Portrait:  b64Portrait,
		}

		monsterDtos = append(monsterDtos, dto)
	}

	data := map[string][]serializedMonsterDto{
		"Monsters": monsterDtos,
	}

	funcMap := template.FuncMap{
		"safe": func(s string) template.URL {
			return template.URL(s)
		},
	}

	t := template.Must(template.New("AllMonsters.html").Funcs(funcMap).ParseFiles(filepath.Join(fileBasePath, "AllMonsters.html")))
	err := t.Execute(w, data)
	if err != nil {
		fmt.Print(err.Error())
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

func createOrUpdateMonster(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer r.Body.Close()

	monsterId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		fmt.Print(err.Error(), monsterId)
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var monster createMonsterDto

	err = decoder.Decode(&monster, r.PostForm)
	if err != nil {
		http.Error(w, "Erreur lors du décodage : "+err.Error(), http.StatusBadRequest)
		return
	}

	var monsterEntity models.Monster
	db := db.GetDbConnection()

	monsterEntity.CampaignID = monster.Campaign
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

	if len(monster.SpecialTraits) > 0 {
		if monster.SpecialTraits[0].Name != "" {
			for _, st := range monster.SpecialTraits {
				newSpecialTrait := models.SpecialTrait{}
				newSpecialTrait.Name = st.Name
				newSpecialTrait.Description = st.Description
				monsterEntity.SpecialTraits = append(monsterEntity.SpecialTraits, newSpecialTrait)
			}
		}
	}

	if len(monster.Actions) > 0 {
		if monster.Actions[0].Name != "" {
			for _, ac := range monster.Actions {
				newAction := models.Action{}
				newAction.Name = ac.Name
				newAction.Description = ac.Description
				newAction.Type = "action"
				monsterEntity.Actions = append(monsterEntity.Actions, newAction)
			}
		}
	}

	if len(monster.Reactions) > 0 {
		if monster.Reactions[0].Name != "" {
			for _, re := range monster.Reactions {
				newReaction := models.Action{}
				newReaction.Name = re.Name
				newReaction.Description = re.Description
				newReaction.Type = "reaction"
				monsterEntity.Actions = append(monsterEntity.Actions, newReaction)
			}
		}
	}

	if len(monster.BonusActions) > 0 {
		if monster.BonusActions[0].Name != "" {
			for _, bac := range monster.BonusActions {
				newBonusAction := models.Action{}
				newBonusAction.Name = bac.Name
				newBonusAction.Description = bac.Description
				newBonusAction.Type = "bonus"
				monsterEntity.Actions = append(monsterEntity.Actions, newBonusAction)
			}
		}
	}

	if len(monster.LegendaryActions) > 0 {
		if monster.LegendaryActions[0].Name != "" {
			for _, lac := range monster.LegendaryActions {
				newLegendaryAction := models.Action{}
				newLegendaryAction.Name = lac.Name
				newLegendaryAction.Description = lac.Description
				newLegendaryAction.Type = "legendary"
				monsterEntity.Actions = append(monsterEntity.Actions, newLegendaryAction)
			}
		}
	}

	file, handler, err := r.FormFile("Portrait")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	monsterEntity.Portrait = fileBytes

	if result := db.Save(&monsterEntity); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}
}

func getEditMonsterTemplate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	skillId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var monster models.Monster
	db := db.GetDbConnection()

	if result := db.Where("ID = ?", skillId).First(&monster); result.Error != nil {
		fmt.Println(result.Error)
	}

	campaigns, err := campaign.GetAllCampaigns()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	data := map[string]interface{}{
		"current":   skillId,
		"next":      skillId + 1,
		"Monster":   monster,
		"Campaigns": campaigns,
	}

	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "EditMonster.html")))
	t.Execute(w, data)
}

func getEditMonsterStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	monsterId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var monster models.Monster
	db := db.GetDbConnection()

	if result := db.Where("ID = ?", monsterId).First(&monster); result.Error != nil {
		fmt.Println(result.Error)
	}

	data := map[string]interface{}{
		"Monster": monster,
	}

	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "/fragments/EditMonster/Stats.html")))
	t.Execute(w, data)
}

func getEditMonsterSkills(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	monsterId, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var monsterSkills []models.SpecialTrait
	db := db.GetDbConnection()

	if result := db.Where("monster_id = ?", monsterId).Find(&monsterSkills); result.Error != nil {
		fmt.Println(result.Error)
	}

	next := len(monsterSkills) + 1

	data := map[string]interface{}{
		"Skills": monsterSkills,
		"next":   next,
	}

	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "/fragments/EditMonster/Skills.html")))
	err = t.Execute(w, data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func getEditMonsterActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	monsterId, err := strconv.Atoi(ps.ByName("id"))
	actionType := ps.ByName("type")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var monsterActions []models.Action
	db := db.GetDbConnection()

	if result := db.Where("monster_id = ? AND type = ?", monsterId, actionType).Find(&monsterActions); result.Error != nil {
		fmt.Println(result.Error)
	}

	next := len(monsterActions) + 1

	data := map[string]interface{}{
		"Actions": monsterActions,
		"next":    next,
		"type":    actionType,
	}

	t := template.Must(template.ParseFiles(filepath.Join(fileBasePath, "/fragments/EditMonster/Actions.html")))
	err = t.Execute(w, data)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func deleteCurrentAbility(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "")
}
