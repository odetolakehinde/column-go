package api

import (
	"context"
	"net/http"

	"github.com/ovalfi/integrations-playground/column/model"
)

// SimulateWireInflow handles the function to simulate a wire transfer inflow
func (c *Call) SimulateWireInflow(ctx context.Context, request model.SimulateWireTransferRequest) error {
	var response struct{}

	err := c.makeRequest(ctx, http.MethodPost, "/simulate/receive-wire", request, response, nil)
	if err != nil {
		return err
	}

	return nil
}

// SimulateIntlWireInflow handles the function to simulate an international wire transfer inflow
func (c *Call) SimulateIntlWireInflow(ctx context.Context, request model.SimulateWireTransferRequest) error {
	var response struct{}

	err := c.makeRequest(ctx, http.MethodPost, "/simulate/receive-international-wire", request, response, nil)
	if err != nil {
		return err
	}

	return nil
}

// SettleWireTransfer forces the submission and settlement of an outgoing Wire transfer, regardless of the operating hours of the Sandbox Fed.
func (c *Call) SettleWireTransfer(ctx context.Context, wireTransferID string) error {
	var response struct{}

	request := struct {
		WireTransferID string `json:"wire_transfer_id"`
	}{
		WireTransferID: wireTransferID,
	}

	err := c.makeRequest(ctx, http.MethodPost, "/simulate/transfers/wire/settle", request, response, nil)
	if err != nil {
		return err
	}

	return nil
}
