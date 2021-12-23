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
