package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	return CampaignTransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	sliceOfTransaction := []CampaignTransactionFormatter{}
	for _, transaction := range transactions {
		formatted := FormatCampaignTransaction(transaction)
		sliceOfTransaction = append(sliceOfTransaction, formatted)
	}
	return sliceOfTransaction
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	imageURL := ""
	for _, image := range transaction.Campaign.CampaignImages {
		if image.IsPrimary {
			imageURL = image.FileName
		}
	}
	return UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign:  CampaignFormatter{transaction.Campaign.Name, imageURL},
	}
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	sliceOfTransaction := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		formatted := FormatUserTransaction(transaction)
		sliceOfTransaction = append(sliceOfTransaction, formatted)
	}
	return sliceOfTransaction
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignId int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	return TransactionFormatter{
		ID:         transaction.ID,
		CampaignId: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentURL: transaction.PaymentURL,
	}
}
