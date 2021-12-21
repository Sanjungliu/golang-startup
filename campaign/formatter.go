package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}
	return CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageURL:         imageURL,
	}

}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	sliceOfCampaigns := []CampaignFormatter{}

	for _, campaign := range campaigns {
		formattedCampaign := FormatCampaign(campaign)
		sliceOfCampaigns = append(sliceOfCampaigns, formattedCampaign)
	}

	return sliceOfCampaigns
}
