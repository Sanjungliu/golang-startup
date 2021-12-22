package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
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
		Slug:             campaign.Slug,
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

type CampaignDetailFormatter struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	ShortDescription string      `json:"short_description"`
	Description      string      `json:"description"`
	ImageURL         string      `json:"image_url"`
	GoalAmount       int         `json:"goal_amount"`
	CurrentAmount    int         `json:"current_amount"`
	UserID           int         `json:"user_id"`
	Slug             string      `json:"slug"`
	Perks            []string    `json:"perks"`
	User             UserData    `json:"user"`
	Images           []ImageData `json:"images"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}

	images := []ImageData{}
	for _, image := range campaign.CampaignImages {
		images = append(images, ImageData{image.FileName, image.IsPrimary})
	}

	perks := strings.Split(campaign.Perks, ", ")
	return CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageURL:         imageURL,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		Perks:            perks,
		User:             UserData{campaign.User.Name, campaign.User.AvatarFileName},
		Images:           images,
	}
}

type UserData struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ImageData struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type CreateCampaignFormatter struct {
}
