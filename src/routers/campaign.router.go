package router

import (
	handler "nearrivers/monster-creator/src/handlers"

	"github.com/julienschmidt/httprouter"
)

func ConfigureCampaignRoutes(r *httprouter.Router) {
	r.GET("/campaign/", handler.GetAllCampaigns)
	r.GET("/campaign/:id", handler.GetOneCampaign)
	r.POST("/campaign/", handler.CreateCampaign)
	r.PUT("/campaign/:id", handler.UpdateCampaign)
	r.DELETE("/campaign/:id", handler.DeleteOneCampaign)
}
