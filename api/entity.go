package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ovalfi/integrations-playground/column/model"
)

// CreatePersonEntity handles the function to create a legal person entity.
func (c *Call) CreatePersonEntity(ctx context.Context, request model.CreatePersonRequest) (*model.PersonEntity, error) {
	ctx, span := c.tracer.Start(ctx, "column.CreatePersonEntity")
	defer span.End()

	response := &model.PersonEntity{}

	err := c.makeRequest(ctx, http.MethodPost, "/entities/person", request, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetPersonEntity handles the function to get a legal person entity by its ID
func (c *Call) GetPersonEntity(ctx context.Context, id string) (*model.PersonEntity, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetPersonEntity")
	defer span.End()

	response := &model.PersonEntity{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/entities/%s", id), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateBusinessEntity handles the function to create a legal business entity.
func (c *Call) CreateBusinessEntity(ctx context.Context, request model.CreateBusinessRequest) (*model.BusinessEntity, error) {
	ctx, span := c.tracer.Start(ctx, "column.CreateBusinessEntity")
	defer span.End()

	response := &model.BusinessEntity{}
	err := c.makeRequest(ctx, http.MethodPost, "/entities/business", request, response, &request.BusinessName)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBusinessEntity handles the function to get a legal business entity by its ID
func (c *Call) GetBusinessEntity(ctx context.Context, id string) (*model.BusinessEntity, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetBusinessEntity")
	defer span.End()

	response := &model.BusinessEntity{}

	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("/entities/%s", id), nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetBusinessEntities handles the function to get all business entities by its ID
func (c *Call) GetBusinessEntities(ctx context.Context) ([]model.BusinessEntity, error) {
	ctx, span := c.tracer.Start(ctx, "column.GetBusinessEntities")
	defer span.End()

	type res struct {
		Entities []model.BusinessEntity `json:"entities"`
	}

	response := &res{}

	err := c.makeRequest(ctx, http.MethodGet, "/entities?limit=100", nil, response, nil)
	if err != nil {
		return nil, err
	}

	return response.Entities, nil
}
