// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/ovalfi/integrations-playground/column/model"
)

// MockRemoteCalls is a mock of RemoteCalls interface.
type MockRemoteCalls struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteCallsMockRecorder
}

// MockRemoteCallsMockRecorder is the mock recorder for MockRemoteCalls.
type MockRemoteCallsMockRecorder struct {
	mock *MockRemoteCalls
}

// NewMockRemoteCalls creates a new mock instance.
func NewMockRemoteCalls(ctrl *gomock.Controller) *MockRemoteCalls {
	mock := &MockRemoteCalls{ctrl: ctrl}
	mock.recorder = &MockRemoteCallsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteCalls) EXPECT() *MockRemoteCallsMockRecorder {
	return m.recorder
}

// BookFxQuote mocks base method.
func (m *MockRemoteCalls) BookFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookFxQuote", ctx, fxQuoteID)
	ret0, _ := ret[0].(*model.FxQuote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookFxQuote indicates an expected call of BookFxQuote.
func (mr *MockRemoteCallsMockRecorder) BookFxQuote(ctx, fxQuoteID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookFxQuote", reflect.TypeOf((*MockRemoteCalls)(nil).BookFxQuote), ctx, fxQuoteID)
}

// CancelFxQuote mocks base method.
func (m *MockRemoteCalls) CancelFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelFxQuote", ctx, fxQuoteID)
	ret0, _ := ret[0].(*model.FxQuote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelFxQuote indicates an expected call of CancelFxQuote.
func (mr *MockRemoteCallsMockRecorder) CancelFxQuote(ctx, fxQuoteID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelFxQuote", reflect.TypeOf((*MockRemoteCalls)(nil).CancelFxQuote), ctx, fxQuoteID)
}

// CancelOutgoingIntlWireTransfer mocks base method.
func (m *MockRemoteCalls) CancelOutgoingIntlWireTransfer(ctx context.Context, request model.CancelIntlWireTransferRequest) (*model.IntlWireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOutgoingIntlWireTransfer", ctx, request)
	ret0, _ := ret[0].(*model.IntlWireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOutgoingIntlWireTransfer indicates an expected call of CancelOutgoingIntlWireTransfer.
func (mr *MockRemoteCallsMockRecorder) CancelOutgoingIntlWireTransfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOutgoingIntlWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).CancelOutgoingIntlWireTransfer), ctx, request)
}

// CreateBankAccount mocks base method.
func (m *MockRemoteCalls) CreateBankAccount(ctx context.Context, request model.CreateBankRequest) (*model.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBankAccount", ctx, request)
	ret0, _ := ret[0].(*model.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBankAccount indicates an expected call of CreateBankAccount.
func (mr *MockRemoteCallsMockRecorder) CreateBankAccount(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBankAccount", reflect.TypeOf((*MockRemoteCalls)(nil).CreateBankAccount), ctx, request)
}

// CreateBankAccountNumber mocks base method.
func (m *MockRemoteCalls) CreateBankAccountNumber(ctx context.Context, request model.CreateBankAccountNoRequest) (*model.BankAccountNumber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBankAccountNumber", ctx, request)
	ret0, _ := ret[0].(*model.BankAccountNumber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBankAccountNumber indicates an expected call of CreateBankAccountNumber.
func (mr *MockRemoteCallsMockRecorder) CreateBankAccountNumber(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBankAccountNumber", reflect.TypeOf((*MockRemoteCalls)(nil).CreateBankAccountNumber), ctx, request)
}

// CreateBusinessEntity mocks base method.
func (m *MockRemoteCalls) CreateBusinessEntity(ctx context.Context, request model.CreateBusinessRequest) (*model.BusinessEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBusinessEntity", ctx, request)
	ret0, _ := ret[0].(*model.BusinessEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBusinessEntity indicates an expected call of CreateBusinessEntity.
func (mr *MockRemoteCallsMockRecorder) CreateBusinessEntity(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBusinessEntity", reflect.TypeOf((*MockRemoteCalls)(nil).CreateBusinessEntity), ctx, request)
}

// CreateCounterparty mocks base method.
func (m *MockRemoteCalls) CreateCounterparty(ctx context.Context, request model.CreateCounterPartyRequest) (*model.CounterParty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCounterparty", ctx, request)
	ret0, _ := ret[0].(*model.CounterParty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCounterparty indicates an expected call of CreateCounterparty.
func (mr *MockRemoteCallsMockRecorder) CreateCounterparty(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCounterparty", reflect.TypeOf((*MockRemoteCalls)(nil).CreateCounterparty), ctx, request)
}

// CreateFxQuote mocks base method.
func (m *MockRemoteCalls) CreateFxQuote(ctx context.Context, request model.CreateFxQuoteRequest) (*model.FxQuote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFxQuote", ctx, request)
	ret0, _ := ret[0].(*model.FxQuote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFxQuote indicates an expected call of CreateFxQuote.
func (mr *MockRemoteCallsMockRecorder) CreateFxQuote(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFxQuote", reflect.TypeOf((*MockRemoteCalls)(nil).CreateFxQuote), ctx, request)
}

// CreatePersonEntity mocks base method.
func (m *MockRemoteCalls) CreatePersonEntity(ctx context.Context, request model.CreatePersonRequest) (*model.PersonEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePersonEntity", ctx, request)
	ret0, _ := ret[0].(*model.PersonEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePersonEntity indicates an expected call of CreatePersonEntity.
func (mr *MockRemoteCallsMockRecorder) CreatePersonEntity(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePersonEntity", reflect.TypeOf((*MockRemoteCalls)(nil).CreatePersonEntity), ctx, request)
}

// GetAllIntlWireTransfer mocks base method.
func (m *MockRemoteCalls) GetAllIntlWireTransfer(ctx context.Context) ([]model.IntlWireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllIntlWireTransfer", ctx)
	ret0, _ := ret[0].([]model.IntlWireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllIntlWireTransfer indicates an expected call of GetAllIntlWireTransfer.
func (mr *MockRemoteCallsMockRecorder) GetAllIntlWireTransfer(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllIntlWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).GetAllIntlWireTransfer), ctx)
}

// GetAllWireTransfer mocks base method.
func (m *MockRemoteCalls) GetAllWireTransfer(ctx context.Context) ([]model.WireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWireTransfer", ctx)
	ret0, _ := ret[0].([]model.WireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWireTransfer indicates an expected call of GetAllWireTransfer.
func (mr *MockRemoteCallsMockRecorder) GetAllWireTransfer(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).GetAllWireTransfer), ctx)
}

// GetBankAccountByID mocks base method.
func (m *MockRemoteCalls) GetBankAccountByID(ctx context.Context, id string) (*model.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBankAccountByID", ctx, id)
	ret0, _ := ret[0].(*model.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBankAccountByID indicates an expected call of GetBankAccountByID.
func (mr *MockRemoteCallsMockRecorder) GetBankAccountByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankAccountByID", reflect.TypeOf((*MockRemoteCalls)(nil).GetBankAccountByID), ctx, id)
}

// GetBankAccounts mocks base method.
func (m *MockRemoteCalls) GetBankAccounts(ctx context.Context) ([]model.BankAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBankAccounts", ctx)
	ret0, _ := ret[0].([]model.BankAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBankAccounts indicates an expected call of GetBankAccounts.
func (mr *MockRemoteCallsMockRecorder) GetBankAccounts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBankAccounts", reflect.TypeOf((*MockRemoteCalls)(nil).GetBankAccounts), ctx)
}

// GetBookTransferByID mocks base method.
func (m *MockRemoteCalls) GetBookTransferByID(ctx context.Context, id string) (*model.BookTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookTransferByID", ctx, id)
	ret0, _ := ret[0].(*model.BookTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookTransferByID indicates an expected call of GetBookTransferByID.
func (mr *MockRemoteCallsMockRecorder) GetBookTransferByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookTransferByID", reflect.TypeOf((*MockRemoteCalls)(nil).GetBookTransferByID), ctx, id)
}

// GetBusinessEntities mocks base method.
func (m *MockRemoteCalls) GetBusinessEntities(ctx context.Context) ([]model.BusinessEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBusinessEntities", ctx)
	ret0, _ := ret[0].([]model.BusinessEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBusinessEntities indicates an expected call of GetBusinessEntities.
func (mr *MockRemoteCallsMockRecorder) GetBusinessEntities(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBusinessEntities", reflect.TypeOf((*MockRemoteCalls)(nil).GetBusinessEntities), ctx)
}

// GetBusinessEntity mocks base method.
func (m *MockRemoteCalls) GetBusinessEntity(ctx context.Context, id string) (*model.BusinessEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBusinessEntity", ctx, id)
	ret0, _ := ret[0].(*model.BusinessEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBusinessEntity indicates an expected call of GetBusinessEntity.
func (mr *MockRemoteCallsMockRecorder) GetBusinessEntity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBusinessEntity", reflect.TypeOf((*MockRemoteCalls)(nil).GetBusinessEntity), ctx, id)
}

// GetFxQuote mocks base method.
func (m *MockRemoteCalls) GetFxQuote(ctx context.Context, fxQuoteID string) (*model.FxQuote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFxQuote", ctx, fxQuoteID)
	ret0, _ := ret[0].(*model.FxQuote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFxQuote indicates an expected call of GetFxQuote.
func (mr *MockRemoteCallsMockRecorder) GetFxQuote(ctx, fxQuoteID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFxQuote", reflect.TypeOf((*MockRemoteCalls)(nil).GetFxQuote), ctx, fxQuoteID)
}

// GetInstitution mocks base method.
func (m *MockRemoteCalls) GetInstitution(ctx context.Context, routingNo string) (*model.Institution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstitution", ctx, routingNo)
	ret0, _ := ret[0].(*model.Institution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstitution indicates an expected call of GetInstitution.
func (mr *MockRemoteCallsMockRecorder) GetInstitution(ctx, routingNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstitution", reflect.TypeOf((*MockRemoteCalls)(nil).GetInstitution), ctx, routingNo)
}

// GetIntlWireTransfer mocks base method.
func (m *MockRemoteCalls) GetIntlWireTransfer(ctx context.Context, swiftTransferID string) (*model.IntlWireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIntlWireTransfer", ctx, swiftTransferID)
	ret0, _ := ret[0].(*model.IntlWireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntlWireTransfer indicates an expected call of GetIntlWireTransfer.
func (mr *MockRemoteCallsMockRecorder) GetIntlWireTransfer(ctx, swiftTransferID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntlWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).GetIntlWireTransfer), ctx, swiftTransferID)
}

// GetPersonEntity mocks base method.
func (m *MockRemoteCalls) GetPersonEntity(ctx context.Context, id string) (*model.PersonEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersonEntity", ctx, id)
	ret0, _ := ret[0].(*model.PersonEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPersonEntity indicates an expected call of GetPersonEntity.
func (mr *MockRemoteCallsMockRecorder) GetPersonEntity(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersonEntity", reflect.TypeOf((*MockRemoteCalls)(nil).GetPersonEntity), ctx, id)
}

// GetWireTransferByID mocks base method.
func (m *MockRemoteCalls) GetWireTransferByID(ctx context.Context, id string) (*model.WireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWireTransferByID", ctx, id)
	ret0, _ := ret[0].(*model.WireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWireTransferByID indicates an expected call of GetWireTransferByID.
func (mr *MockRemoteCallsMockRecorder) GetWireTransferByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWireTransferByID", reflect.TypeOf((*MockRemoteCalls)(nil).GetWireTransferByID), ctx, id)
}

// InitiateBookTransfer mocks base method.
func (m *MockRemoteCalls) InitiateBookTransfer(ctx context.Context, request model.CreateBookTransfer) (*model.BookTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiateBookTransfer", ctx, request)
	ret0, _ := ret[0].(*model.BookTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitiateBookTransfer indicates an expected call of InitiateBookTransfer.
func (mr *MockRemoteCallsMockRecorder) InitiateBookTransfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiateBookTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).InitiateBookTransfer), ctx, request)
}

// InitiateIntlWireTransfer mocks base method.
func (m *MockRemoteCalls) InitiateIntlWireTransfer(ctx context.Context, request model.CreateWireTransferRequest) (*model.IntlWireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiateIntlWireTransfer", ctx, request)
	ret0, _ := ret[0].(*model.IntlWireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitiateIntlWireTransfer indicates an expected call of InitiateIntlWireTransfer.
func (mr *MockRemoteCallsMockRecorder) InitiateIntlWireTransfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiateIntlWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).InitiateIntlWireTransfer), ctx, request)
}

// InitiateWireTransfer mocks base method.
func (m *MockRemoteCalls) InitiateWireTransfer(ctx context.Context, request model.CreateWireTransferRequest) (*model.WireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitiateWireTransfer", ctx, request)
	ret0, _ := ret[0].(*model.WireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitiateWireTransfer indicates an expected call of InitiateWireTransfer.
func (mr *MockRemoteCallsMockRecorder) InitiateWireTransfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitiateWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).InitiateWireTransfer), ctx, request)
}

// ListBankAccountNumbers mocks base method.
func (m *MockRemoteCalls) ListBankAccountNumbers(ctx context.Context, accountID string) ([]model.BankAccountNumber, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBankAccountNumbers", ctx, accountID)
	ret0, _ := ret[0].([]model.BankAccountNumber)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBankAccountNumbers indicates an expected call of ListBankAccountNumbers.
func (mr *MockRemoteCallsMockRecorder) ListBankAccountNumbers(ctx, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBankAccountNumbers", reflect.TypeOf((*MockRemoteCalls)(nil).ListBankAccountNumbers), ctx, accountID)
}

// ListCounterparties mocks base method.
func (m *MockRemoteCalls) ListCounterparties(ctx context.Context) ([]model.CounterParty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCounterparties", ctx)
	ret0, _ := ret[0].([]model.CounterParty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCounterparties indicates an expected call of ListCounterparties.
func (mr *MockRemoteCallsMockRecorder) ListCounterparties(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCounterparties", reflect.TypeOf((*MockRemoteCalls)(nil).ListCounterparties), ctx)
}

// ReverseIncomingIntlWireTransfer mocks base method.
func (m *MockRemoteCalls) ReverseIncomingIntlWireTransfer(ctx context.Context, request model.ReverseIntlWireTransferRequest) (*model.IntlWireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseIncomingIntlWireTransfer", ctx, request)
	ret0, _ := ret[0].(*model.IntlWireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReverseIncomingIntlWireTransfer indicates an expected call of ReverseIncomingIntlWireTransfer.
func (mr *MockRemoteCallsMockRecorder) ReverseIncomingIntlWireTransfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseIncomingIntlWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).ReverseIncomingIntlWireTransfer), ctx, request)
}

// ReverseIncomingWireTransfer mocks base method.
func (m *MockRemoteCalls) ReverseIncomingWireTransfer(ctx context.Context, request model.ReverseWireTransferRequest) (*model.WireTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseIncomingWireTransfer", ctx, request)
	ret0, _ := ret[0].(*model.WireTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReverseIncomingWireTransfer indicates an expected call of ReverseIncomingWireTransfer.
func (mr *MockRemoteCallsMockRecorder) ReverseIncomingWireTransfer(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseIncomingWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).ReverseIncomingWireTransfer), ctx, request)
}

// RunInSandboxMode mocks base method.
func (m *MockRemoteCalls) RunInSandboxMode() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunInSandboxMode")
}

// RunInSandboxMode indicates an expected call of RunInSandboxMode.
func (mr *MockRemoteCallsMockRecorder) RunInSandboxMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunInSandboxMode", reflect.TypeOf((*MockRemoteCalls)(nil).RunInSandboxMode))
}

// SettleWireTransfer mocks base method.
func (m *MockRemoteCalls) SettleWireTransfer(ctx context.Context, wireTransferID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettleWireTransfer", ctx, wireTransferID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SettleWireTransfer indicates an expected call of SettleWireTransfer.
func (mr *MockRemoteCallsMockRecorder) SettleWireTransfer(ctx, wireTransferID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettleWireTransfer", reflect.TypeOf((*MockRemoteCalls)(nil).SettleWireTransfer), ctx, wireTransferID)
}

// SimulateIntlWireInflow mocks base method.
func (m *MockRemoteCalls) SimulateIntlWireInflow(ctx context.Context, request model.SimulateWireTransferRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateIntlWireInflow", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SimulateIntlWireInflow indicates an expected call of SimulateIntlWireInflow.
func (mr *MockRemoteCallsMockRecorder) SimulateIntlWireInflow(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateIntlWireInflow", reflect.TypeOf((*MockRemoteCalls)(nil).SimulateIntlWireInflow), ctx, request)
}

// SimulateWireInflow mocks base method.
func (m *MockRemoteCalls) SimulateWireInflow(ctx context.Context, request model.SimulateWireTransferRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateWireInflow", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// SimulateWireInflow indicates an expected call of SimulateWireInflow.
func (mr *MockRemoteCallsMockRecorder) SimulateWireInflow(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateWireInflow", reflect.TypeOf((*MockRemoteCalls)(nil).SimulateWireInflow), ctx, request)
}
