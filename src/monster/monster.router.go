package monster

import "github.com/julienschmidt/httprouter"

func ConfigureRoutes(r *httprouter.Router) {
	r.GET("/monster/list", getAllMonsters)
	r.GET("/monster/new", getNewMonsterTemplate)
	r.GET("/monster/stats/new", getNewMonsterStats)
	r.GET("/monster/skills/new/:id", getNewMonsterSkills)
	r.GET("/monster/actions/new/:id/:type", getNewMonsterActions)

	r.GET("/monster/edit/:id", editMonsterTemplate)
	r.GET("/monster/stats/edit/:id", editMonsterStats)
	// r.GET("/monster/skills/edit")
	// r.GET("/monster/actions/edit")

	r.POST("/monster/new", createMonster)

	r.DELETE("/monster/skills/new", deleteCurrentAbility)
	r.DELETE("/monster/actions/new", deleteCurrentAbility)
}
