package transaction

import "github.com/Sanjungliu/golang-startup/user"

type GetTransactionByCampaignIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
