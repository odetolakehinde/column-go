// Package api defines implementations of endpoints and calls
package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/ovalfi/integrations-playground/column/model"
)

// RemoteCalls abstracted definition of supported functions
//
//go:generate mockgen -source api.go -destination ./mock/mock_remote_calls.go -package mock RemoteCalls
type RemoteCalls interface {
	RunInSandboxMode()

	CreatePersonEntity(ctx context.Context, request model.CreatePersonRequest) (*model.PersonEntity, error)
	CreateBusinessEntity(ctx context.Context, request model.CreateBusinessRequest) (*model.BusinessEntity, error)
	GetPersonEntity(ctx context.Context, id string) (*model.PersonEntity, error)
	GetBusinessEntity(ctx context.Context, id string) (*model.BusinessEntity, error)
	GetBusinessEntities(ctx context.Context) ([]model.BusinessEntity, error)

	CreateBankAccount(ctx context.Context, request model.CreateBankRequest) (*model.BankAccount, error)
	GetBankAccountByID(ctx context.Context, id string) (*model.BankAccount, error)
	CreateBankAccountNumber(ctx context.Context, request model.CreateBankAccountNoRequest) (*model.BankAccountNumber, error)
	ListBankAccountNumbers(ctx context.Context, accountID string) ([]model.BankAccountNumber, error)
	GetBankAccounts(ctx context.Context) ([]model.BankAccount, error)

	InitiateWireTransfer(ctx context.Context, request model.CreateWireTransferRequest) (*model.WireTransfer, error)
	GetWireTransferByID(ctx context.Context, id string) (*model.WireTransfer, error)
	ReverseIncomingWireTransfer(ctx context.Context, request model.ReverseWireTransferRequest) (*model.WireTransfer, error)
	ReverseIncomingIntlWireTransfer(ctx context.Context, request model.ReverseIntlWireTransferRequest) (*model.IntlWireTransfer, error)
	CancelOutgoingIntlWireTransfer(ctx context.Context, request model.CancelIntlWireTransferRequest) (*model.IntlWireTransfer, error)
	InitiateIntlWireTransfer(ctx context.Context, request model.CreateWireTransferRequest) (*model.IntlWireTransfer, error)
	GetIntlWireTransfer(ctx context.Context, swiftTransferID string) (*model.IntlWireTransfer, error)
	CreateCounterparty(ctx context.Context, request model.CreateCounterPartyRequest) (*model.CounterParty, error)
	ListCounterparties(ctx context.Context) ([]model.CounterParty, error)
	GetInstitution(ctx context.Context, routingNo string) (*model.Institution, error)
	InitiateBookTransfer(ctx context.Context, request model.CreateBookTransfer) (*model.BookTransfer, error)
	GetBookTransferByID(ctx context.Context, id string) (*model.BookTransfer, error)
	GetAllIntlWireTransfer(ctx context.Context) ([]model.IntlWireTransfer, error)
	GetAllWireTransfer(ctx context.Context) ([]model.WireTransfer, error)

	SimulateWireInflow(ctx context.Context, request model.SimulateWireTransferRequest) error
	SimulateIntlWireInflow(ctx context.Context, request model.SimulateWireTransferRequest) error
	SettleWireTransfer(ctx context.Context, wireTransferID string) error

	CreateFxQuote(ctx context.Context, request model.CreateFxQuoteRequest) (*model.FxQuote, error)
	GetFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error)
	BookFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error)
	CancelFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error)
}

// Call object
type Call struct {
	client          *resty.Client
	logger          zerolog.Logger
	sandboxMode     bool
	baseURL         string
	proxyURL        string
	tracer          trace.Tracer
	isAuthenticated bool
}

// New initialises the object Call
func New(z *zerolog.Logger, bURL, apiKey, proxyURL string) RemoteCalls {
	c := createInstrumentedRestyClientWithProxy(proxyURL)
	call := &Call{
		client:   c,
		logger:   z.With().Str("sdk", "Column").Logger(),
		baseURL:  bURL,
		proxyURL: proxyURL,
		tracer:   otel.Tracer("column"),
	}
	c.SetBasicAuth("", apiKey)
	c.SetHeader(model.HeaderContentTypeKey, model.HeaderContentTypeValue)
	c.SetHeader(model.HeaderContentAcceptKey, model.HeaderContentTypeValue)

	call.authenticate()

	return RemoteCalls(call)
}

// RunInSandboxMode this forces Call functionalities to run in sandbox mode for relevant logic/API consumption
func (c *Call) RunInSandboxMode() {
	c.client.SetDebug(true)
	c.client.EnableTrace()
	c.sandboxMode = true
}

func (c *Call) authenticate() {
	// This function calls a valid get request endpoint to ensure the apiKey or apiSecret is valid, since the provider
	// provides no way to validate this. When the request returns status 200 or 204, c.isAuthenticated is set to true.
	// This will block unauthenticated request from going through.
	response, err := c.client.R().
		SetContext(context.Background()).
		SetDebug(false).
		Get(fmt.Sprintf("%s/bank-accounts?limit=5", c.baseURL))
	if err != nil {
		c.logger.Err(err).Msg("cannot set api as authenticated")
		return
	}
	if response.StatusCode() == http.StatusOK || response.StatusCode() == http.StatusNoContent {
		c.isAuthenticated = true
	}
}

func createInstrumentedRestyClientWithProxy(proxyURL string) *resty.Client {
	httpClient := &http.Client{}

	if len(proxyURL) > 0 {
		// Create a new base *http.Transport
		baseTransport := &http.Transport{}

		// Parse the proxy URL
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			return resty.New()
		}

		// Set the proxy on the base transport
		baseTransport.Proxy = http.ProxyURL(proxy)

		// Wrap the modified base transport with OpenTelemetry instrumentation
		instrumentedTransport := otelhttp.NewTransport(baseTransport)

		// Create a new http.Client using the instrumented transport
		httpClient = &http.Client{
			Transport: instrumentedTransport,
		}
	}

	// return a new resty client with the custom http.Client
	return resty.NewWithClient(httpClient)
}
