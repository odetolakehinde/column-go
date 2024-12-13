package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/integrations-playground/column/model"
)

// CreateBankAccount handles the function to create a new bank account under an entity.
func (c *Call) CreateBankAccount(ctx context.Context, request model.CreateBankRequest) (*model.BankAccount, error) {
	ctx, span := c.tracer.Start(ctx, "column.CreateBankAccount")
	defer span.End()

	response := &model.BankAccount{}
	err := c.makeRequest(ctx, http.MethodPost, "/bank-accounts", request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBankAccountByID handles the function to get a bank account under an entity by its ID
func (c *Call) GetBankAccountByID(ctx context.Context, id string) (*model.BankAccount, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetBankAccountByID")
	defer span.End()

	response := &model.BankAccount{}
	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/bank-accounts/%s", id), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBankAccounts handles the function to get bank accounts
func (c *Call) GetBankAccounts(ctx context.Context) ([]model.BankAccount, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetBankAccounts")
	defer span.End()

	type res struct {
		BankAccounts []model.BankAccount `json:"bank_accounts"`
	}

	response := &res{}
	err := c.makeRequest(ctx, http.MethodGet, "/bank-accounts?limit=100", nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response.BankAccounts, nil
}

// CreateBankAccountNumber handles the function to create a new account number that points to the associated bank account.
func (c *Call) CreateBankAccountNumber(ctx context.Context, request model.CreateBankAccountNoRequest) (*model.BankAccountNumber, error) {
	ctx, span := c.tracer.Start(ctx, "column.CreateBankAccountNumber")
	defer span.End()

	response := &model.BankAccountNumber{}
	err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("/bank-accounts/%s/account-numbers", request.BankAccountID), request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ListBankAccountNumbers handles the function to list all the account numbers that point to a specific bank account
func (c *Call) ListBankAccountNumbers(ctx context.Context, accountID string) ([]model.BankAccountNumber, error) {
	ctx, span := c.tracer.Start(ctx, "column.ListBankAccountNumbers")
	defer span.End()

	response := &struct {
		AccountNumbers []model.BankAccountNumber `json:"account_numbers"`
		HasMore        bool                      `json:"has_more"`
	}{}

	path := fmt.Sprintf("/bank-accounts/%s/account-numbers", accountID)

	err := c.makeRequest(ctx, http.MethodGet, path, nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response.AccountNumbers, err
}
