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
	var sliceOfTransaction []CampaignTransactionFormatter
	for _, transaction := range transactions {
		formatted := FormatCampaignTransaction(transaction)
		sliceOfTransaction = append(sliceOfTransaction, formatted)
	}
	return sliceOfTransaction
}
