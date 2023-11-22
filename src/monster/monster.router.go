package monster

import "github.com/julienschmidt/httprouter"

func ConfigureRoutes(r *httprouter.Router) {
	r.GET("/monster/list", getAllMonsters)
	r.GET("/monster/new", getNewMonsterTemplate)
	r.GET("/monster/new/stats", getNewMonsterStats)
	r.GET("/monster/new/skills/:id", getNewMonsterSkills)
	r.GET("/monster/new/actions/:id/:type", getNewMonsterActions)
	r.GET("/monster/edit/:id", editMonsterTemplate)

	r.POST("/monster/new", createMonster)

	r.DELETE("/monster/new/skills", deleteCurrentAbility)
	r.DELETE("/monster/new/actions", deleteCurrentAbility)
}
