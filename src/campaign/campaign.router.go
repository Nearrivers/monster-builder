package campaign

import (
	"github.com/julienschmidt/httprouter"
)

func ConfigureRoutes(r *httprouter.Router) {
	r.GET("/campaign/list", GetAllCampaigns)
	r.GET("/campaign/list/:id", GetOneCampaign)
	r.GET("/campaign/monster/new", GetAllCampaignsSelect)

	r.POST("/campaign", CreateCampaign)

	r.PUT("/campaign/:id", UpdateCampaign)

	r.DELETE("/campaign/:id", DeleteOneCampaign)
}
