package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/ovalfi/integrations-playground/column/model"
)

// makeRequest is a super function that handles all requests to Column SDK
func (c *Call) makeRequest(ctx context.Context, method, path string, body, expectedRes interface{}, idempotencyKey *string) error {
	endpoint := fmt.Sprintf("%s%s", c.baseURL, path)

	log := c.logger.With().Str("method", method).Str("endpoint", endpoint).Logger()
	log.Info().Msg("starting...")
	log.Info().Interface(model.LogStrRequest, body).Msg("request")
	defer log.Info().Msg("done...")

	if !c.isAuthenticated {
		err := errors.New("not authenticated, aborting request")
		log.Err(err).Msg("request aborted")
		return err
	}

	var (
		err    error
		res    *resty.Response
		resp   = &model.Response{}
		errRes model.ErrorResponse
	)

	client := c.client.
		R().
		SetContext(ctx).
		SetResult(&expectedRes).
		SetError(&errRes)

	if body != nil {
		client.SetBody(body)
	}

	if idempotencyKey != nil {
		c.client.SetHeader(model.IdempotencyKey, *idempotencyKey)
	}

	switch method {
	case http.MethodGet:
		res, err = client.Get(endpoint)
	case http.MethodPost:
		res, err = client.Post(endpoint)
	case http.MethodPut:
		res, err = client.Put(endpoint)
	case http.MethodPatch:
		res, err = client.Patch(endpoint)
	default:
		err = errors.New("invalid method")
		log.Err(err).Str("method", method).Msg("invalid method passed")
		return err
	}

	if err != nil {
		log.Err(err).Msg("something went wrong")
		if res != nil {
			log.Err(err).Str("error_code", fmt.Sprintf("%d", res.StatusCode())).Msg(string(res.Body()))
		}
		return err
	}

	// check to ensure error is not empty
	if errRes != (model.ErrorResponse{}) {
		err = errors.New(c.translateColumnError(errRes))
		log.Err(err).Msg("error while making request")
		return err
	}

	log.Info().Interface(model.LogStrResponse, resp).Msg("response")

	return nil
}

func (c *Call) translateColumnError(errRes model.ErrorResponse) string {
	log := c.logger.With().Interface("payload", errRes).Str("method", "translateColumnError").Logger()

	errCode, found := model.ErrorDictionary[errRes.Code]
	if !found {
		log.Warn().Msg("Cannot find error code in dictionary. Returning original message.")
		return errRes.Message
	}

	// Handle known error codes with specific formatting
	switch errCode {
	case string(model.ErrorCodeInvalidStringCharacter):
		// Assert details to a map for easier access
		details, ok := errRes.Details.(map[string]interface{})
		if !ok {
			log.Info().Msg("Invalid error details format")
			return errRes.Message
		}

		// Extract specific fields from the details map
		field, fieldOk := details["field"].(string)
		invalidValue, valueOk := details["invalid_value"].(string)

		if !fieldOk || !valueOk {
			log.Info().Msg("missing or invalid fields in error details. Returning original message.")
			return errRes.Message
		}

		// Construct the custom error message
		errMessage := fmt.Sprintf("invalid string character for field %s; %s", field, invalidValue)
		return errMessage

	case string(model.ErrorCodeMandatoryParameterMissing):
		errMessage := fmt.Sprintf("Mandatory parameter missing; %s", errRes.Message)
		return errMessage

	default:
		log.Info().Msg("Unhandled error code. Returning original message.")
		return errRes.Message
	}
}
