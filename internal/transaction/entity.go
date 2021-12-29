package transaction

import (
	"time"

	"github.com/Sanjungliu/golang-startup/internal/campaign"
	"github.com/Sanjungliu/golang-startup/internal/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
