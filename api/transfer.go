package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ovalfi/integrations-playground/column/model"
)

// InitiateWireTransfer handles the function to create a wire transfer
func (c *Call) InitiateWireTransfer(ctx context.Context, request model.CreateWireTransferRequest) (*model.WireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.InitiateWireTransfer")
	defer span.End()

	response := &model.WireTransfer{}

	if len(request.Reference) == 0 {
		return nil, errors.New("reference not provided")
	}

	err := c.makeRequest(ctx, http.MethodPost, "/transfers/wire", request, response, &request.Reference)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetWireTransferByID handles the function to get a wire transfer by its ID
func (c *Call) GetWireTransferByID(ctx context.Context, id string) (*model.WireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetWireTransferByID")
	defer span.End()

	response := &model.WireTransfer{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/transfers/wire/%s?expand=raw_message", id), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllWireTransfer handles the function to get all wire transfers
func (c *Call) GetAllWireTransfer(ctx context.Context) ([]model.WireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetAllWireTransfer")
	defer span.End()

	type res struct {
		Transfers []model.WireTransfer `json:"transfers"`
	}

	response := &res{}
	err := c.makeRequest(ctx, http.MethodGet, "/transfers/wire", nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response.Transfers, nil
}

// ReverseIncomingWireTransfer handles the function to reverse an incoming wire transfer received in error
func (c *Call) ReverseIncomingWireTransfer(ctx context.Context, request model.ReverseWireTransferRequest) (*model.WireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.ReverseIncomingWireTransfer")
	defer span.End()

	response := &model.WireTransfer{}

	err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("/transfers/wire/%s/reverse", request.ID), request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReverseIncomingIntlWireTransfer handles the function to reverse an incoming int'l wire transfer received in error
func (c *Call) ReverseIncomingIntlWireTransfer(ctx context.Context, request model.ReverseIntlWireTransferRequest) (*model.IntlWireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.ReverseIncomingIntlWireTransfer")
	defer span.End()

	response := &model.IntlWireTransfer{}

	err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("/transfers/international-wire/%s/return", request.ID), request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// InitiateIntlWireTransfer handles the function to initiate an int'l wire transfer
func (c *Call) InitiateIntlWireTransfer(ctx context.Context, request model.CreateWireTransferRequest) (*model.IntlWireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.InitiateIntlWireTransfer")
	defer span.End()

	response := &model.IntlWireTransfer{}

	if len(request.Reference) == 0 {
		return nil, errors.New("reference not provided")
	}

	err := c.makeRequest(ctx, http.MethodPost, "/transfers/international-wire", request, response, &request.Reference)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetIntlWireTransfer handles the function to get an int'l wire transfer
func (c *Call) GetIntlWireTransfer(ctx context.Context, swiftTransferID string) (*model.IntlWireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetIntlWireTransfer")
	defer span.End()

	response := &model.IntlWireTransfer{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/transfers/international-wire/%s?expand=raw_message", swiftTransferID), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllIntlWireTransfer handles the function to get all int'l wire transfers
func (c *Call) GetAllIntlWireTransfer(ctx context.Context) ([]model.IntlWireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetAllIntlWireTransfer")
	defer span.End()

	type res struct {
		Transfers []model.IntlWireTransfer `json:"transfers"`
	}

	response := &res{}
	err := c.makeRequest(ctx, http.MethodGet, "/transfers/international-wire", nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response.Transfers, nil
}

// CancelOutgoingIntlWireTransfer handles the function to cancel an outgoing int'l wire transfer received in error
func (c *Call) CancelOutgoingIntlWireTransfer(ctx context.Context, request model.CancelIntlWireTransferRequest) (*model.IntlWireTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.CancelOutgoingIntlWireTransfer")
	defer span.End()

	response := &model.IntlWireTransfer{}

	err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("/transfers/international-wire/%s/cancel", request.ID), request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateCounterparty handles the function to create a new counterparty
func (c *Call) CreateCounterparty(ctx context.Context, request model.CreateCounterPartyRequest) (*model.CounterParty, error) {
	ctx, span := c.tracer.Start(ctx, "column.CreateCounterparty")
	defer span.End()

	response := &model.CounterParty{}

	err := c.makeRequest(ctx, http.MethodPost, "/counterparties", request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ListCounterparties handles the function to get all counterparties information
func (c *Call) ListCounterparties(ctx context.Context) ([]model.CounterParty, error) {
	ctx, span := c.tracer.Start(ctx, "column.ListCounterparties")
	defer span.End()

	type res struct {
		Counterparties []model.CounterParty `json:"counterparties"`
	}

	response := &res{}
	err := c.makeRequest(ctx, http.MethodGet, "/counterparties?limit=100", nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response.Counterparties, nil
}

// GetInstitution handles the function to get an institution/bank details by its routing number
func (c *Call) GetInstitution(ctx context.Context, routingNo string) (*model.Institution, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetInstitution")
	defer span.End()

	response := &model.Institution{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/institutions/%s", routingNo), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// InitiateBookTransfer handles the function to create a book transfer
func (c *Call) InitiateBookTransfer(ctx context.Context, request model.CreateBookTransfer) (*model.BookTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.InitiateBookTransfer")
	defer span.End()

	response := &model.BookTransfer{}

	if len(request.Reference) == 0 {
		return nil, errors.New("reference not provided")
	}

	err := c.makeRequest(ctx, http.MethodPost, "/transfers/book", request, response, &request.Reference)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBookTransferByID handles the function to get a book transfer by its ID
func (c *Call) GetBookTransferByID(ctx context.Context, id string) (*model.BookTransfer, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetBookTransferByID")
	defer span.End()

	response := &model.BookTransfer{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/transfers/book/%s", id), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateFxQuote handles the function to create a fx quote
// see more details - https://column.com/docs/api/#international-wire/fx-create
func (c *Call) CreateFxQuote(ctx context.Context, request model.CreateFxQuoteRequest) (*model.FxQuote, error) {
	ctx, span := c.tracer.Start(ctx, "column.CreateFXQuote")
	defer span.End()

	response := &model.FxQuote{}

	err := c.makeRequest(ctx, http.MethodPost, "/transfers/international-wire/fx-rate", request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetFxQuote handles the function to get a single foreign exchange quote by its ID.
// see more details - https://column.com/docs/api/#international-wire/fx-get
func (c *Call) GetFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetFxQuote")
	defer span.End()

	response := &model.FxQuote{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/transfers/international-wire/fx-rate/%s", fxQuoteID), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// BookFxQuote handles the function to book a single foreign exchange quote by its ID. If the quote has already been booked before, it will be returned.
// see more details - https://column.com/docs/api/#international-wire/fx-book
func (c *Call) BookFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error) {
	ctx, span := c.tracer.Start(ctx, "column.BookFXQuote")
	defer span.End()

	response := &model.FxQuote{}

	err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("/transfers/international-wire/fx-rate/%s/book", fxQuoteID), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CancelFxQuote handles the function to cancel a single foreign exchange quote by its ID. If the quote has already been cancelled before, it will be returned.
// see more details - https://column.com/docs/api/#international-wire/fx-cancel
func (c *Call) CancelFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error) {
	ctx, span := c.tracer.Start(ctx, "column.CancelFXQuote")
	defer span.End()

	response := &model.FxQuote{}

	err := c.makeRequest(ctx, http.MethodPost, fmt.Sprintf("/transfers/international-wire/fx-rate/%s/cancel", fxQuoteID), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}
