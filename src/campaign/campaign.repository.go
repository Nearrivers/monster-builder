package campaign

import (
	"errors"
	"nearrivers/monster-creator/src/db"
	"nearrivers/monster-creator/src/models"
)

func GetAllCampaigns() ([]models.Campaign, error) {
	var campaigns []models.Campaign
	db := db.GetDbConnection()

	if result := db.Find(&campaigns); result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}

	if len(campaigns) == 0 {
		return nil, errors.New("aucune campagne trouvée")
	}

	return campaigns, nil
}
