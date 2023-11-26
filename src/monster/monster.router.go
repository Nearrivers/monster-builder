package monster

import "github.com/julienschmidt/httprouter"

func ConfigureRoutes(r *httprouter.Router) {
	r.GET("/monster/list", getAllMonsters)
	r.GET("/monster/new", getNewMonsterTemplate)
	r.GET("/monster/stats/new", getNewMonsterStats)
	r.GET("/monster/skills/new/:id", getNewMonsterSkills)
	r.GET("/monster/actions/new/:id/:type", getNewMonsterActions)

	r.GET("/monster/edit/:id", getEditMonsterTemplate)
	r.GET("/monster/stats/edit/:id", getEditMonsterStats)
	r.GET("/monster/skills/edit/:id", getEditMonsterSkills)
	r.GET("/monster/actions/edit/:id/:type", getEditMonsterActions)

	r.POST("/monster/new", createOrUpdateMonster)

	r.PUT("/monster/edit/:id", createOrUpdateMonster)

	r.DELETE("/monster/skills/new", deleteCurrentAbility)
	r.DELETE("/monster/actions/new", deleteCurrentAbility)
}
