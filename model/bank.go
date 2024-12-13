package model

type (
	// CreateBankRequest schema
	CreateBankRequest struct {
		EntityID    string `json:"entity_id"`
		Description string `json:"description"`
	}

	// CreateBankAccountNoRequest schema
	CreateBankAccountNoRequest struct {
		BankAccountID string `json:"-"`
		Description   string `json:"description"`
	}

	// BankAccount schema
	BankAccount struct {
		ID                        string   `json:"id"`
		Type                      string   `json:"type"`
		Balances                  Balances `json:"balances"`
		Description               string   `json:"description"`
		CurrencyCode              string   `json:"currency_code"`
		DefaultAccountNumber      string   `json:"default_account_number"`
		DefaultAccountNumberID    string   `json:"default_account_number_id"`
		IsOverdraftable           bool     `json:"is_overdraftable"`
		OverdraftReserveAccountID string   `json:"overdraft_reserve_account_id"`
		Owners                    []string `json:"owners"`
		RoutingNumber             string   `json:"routing_number"`
		CreatedAt                 string   `json:"created_at"`
	}

	// BankAccountNumber account number is the child of a bank account
	BankAccountNumber struct {
		ID            string `json:"id"`
		CreatedAt     string `json:"created_at"`
		BankAccountID string `json:"bank_account_id"`
		Description   string `json:"description"`
		RoutingNumber string `json:"routing_number"`
		AccountNumber string `json:"account_number"`
	}

	// Balances schema
	Balances struct {
		AvailableAmount int `json:"available_amount"`
		HoldingAmount   int `json:"holding_amount"`
		PendingAmount   int `json:"pending_amount"`
		LockedAmount    int `json:"locked_amount"`
	}
)
