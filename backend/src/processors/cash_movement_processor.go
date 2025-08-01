package processors

import (
	"strings"

	"github.com/username/taxfolio/backend/src/models"
)

// cashMovementProcessor implements the CashMovementProcessor interface.
type cashMovementProcessor struct{}

// NewCashMovementProcessor creates a new instance of CashMovementProcessor.
func NewCashMovementProcessor() CashMovementProcessor {
	return &cashMovementProcessor{}
}

// Process identifies cash deposits and withdrawals from the list of processed transactions.
func (p *cashMovementProcessor) Process(transactions []models.ProcessedTransaction) []models.CashMovement {
	var cashMovements []models.CashMovement

	for _, tx := range transactions {
		// Check for cash deposits
		if strings.ToLower(tx.TransactionType) == "cash" && strings.ToLower(tx.TransactionSubType) == "deposit" {
			movement := models.CashMovement{
				Date:     tx.Date,
				Type:     "deposit", // Currently only handling deposits
				Amount:   tx.Amount,
				Currency: tx.Currency,
			}
			cashMovements = append(cashMovements, movement)
		}
		// TODO: Add logic for withdrawals if needed, e.g., check for a specific OrderType or Description
	}

	// TODO: Consider sorting cashMovements by date if necessary

	return cashMovements
}
