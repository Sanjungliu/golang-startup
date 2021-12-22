package handler

import (
	"net/http"
	"strconv"

	"github.com/Sanjungliu/golang-startup/campaign"
	"github.com/Sanjungliu/golang-startup/helper"
	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *CampaignHandler {
	return &CampaignHandler{service}
}

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedCampaigns := campaign.FormatCampaigns(campaigns)

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formattedCampaigns)
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetail

	if err := c.ShouldBindUri(&input); err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignData, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatted := campaign.FormatCampaignDetail(campaignData)
	response := helper.APIResponse("Succeed get campaign detail", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}
