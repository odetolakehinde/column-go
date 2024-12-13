package model

type (
	// SimulateWireTransferRequest is the schema to simulate receiving a wire or int'l wire transfer.
	SimulateWireTransferRequest struct {
		Amount                   int64  `json:"amount"`
		CurrencyCode             string `json:"currency_code"`
		DestinationAccountNumber string `json:"destination_account_number_id"`
	}
)
