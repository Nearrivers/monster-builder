package campaign

import "github.com/julienschmidt/httprouter"

func ConfigureCampaignRoutes(r *httprouter.Router) {
	r.GET("/campaign/", GetAllCampaigns)
	r.GET("/campaign/:id", GetOneCampaign)
	r.POST("/campaign/", CreateCampaign)
	r.PUT("/campaign/:id", UpdateCampaign)
	r.DELETE("/campaign/:id", DeleteOneCampaign)
}
