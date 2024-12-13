package model

import (
	"fmt"
	"time"
)

type (
	// CreateWireTransferRequest to initiate a wire transfer
	CreateWireTransferRequest struct {
		Amount               int64           `json:"amount"`
		CurrencyCode         string          `json:"currency_code"`
		Description          string          `json:"description"`
		BankAccountID        string          `json:"bank_account_id,omitempty"`
		AccountNumberID      string          `json:"account_number_id,omitempty"`
		CounterpartyID       string          `json:"counterparty_id"`
		MsgToBeneficiaryBank string          `json:"message_to_beneficiary_bank,omitempty"`
		IntermediaryBank     string          `json:"intermediary_bank,omitempty"`
		RemittanceInfo       *RemittanceInfo `json:"remittance_info,omitempty"`
		FxQuoteID            string          `json:"fx_quote_id,omitempty"`
		Reference            string          `json:"-"`
	}

	// RemittanceInfo schema
	RemittanceInfo struct {
		GeneralInfo          string `json:"general_info,omitempty"`
		BeneficiaryReference string `json:"beneficiary_reference,omitempty"`
	}

	// ReverseWireTransferRequest to reverse an incoming wire transfer
	ReverseWireTransferRequest struct {
		ID          string       `json:"-"`
		Reason      RejectReason `json:"reason"`
		Description string       `json:"description"`
	}

	// ReverseIntlWireTransferRequest to reverse an incoming int'l wire transfer
	ReverseIntlWireTransferRequest struct {
		ID                  string       `json:"-"`
		ReturnReason        RejectReason `json:"return_reason"`
		MsgToOriginatorBank string       `json:"message_to_originator_bank,omitempty"`
	}

	// CancelIntlWireTransferRequest to cancel an outgoing int'l wire transfer
	CancelIntlWireTransferRequest struct {
		ID           string       `json:"-"`
		CancelReason CancelReason `json:"cancellation_reason"`
	}

	// WireTransfer schema
	WireTransfer struct {
		AccountNumberID                         string   `json:"account_number_id"`
		AllowOverdraft                          bool     `json:"allow_overdraft"`
		Amount                                  int      `json:"amount"`
		BankAccountID                           string   `json:"bank_account_id"`
		BeneficiaryAccountNumber                string   `json:"beneficiary_account_number"`
		BeneficiaryName                         string   `json:"beneficiary_name"`
		BeneficiaryReference                    string   `json:"beneficiary_reference"`
		BusinessFunctionCode                    string   `json:"business_function_code"`
		CompletedAt                             string   `json:"completed_at"`
		CounterpartyID                          string   `json:"counterparty_id"`
		CreatedAt                               string   `json:"created_at"`
		CurrencyCode                            string   `json:"currency_code"`
		Description                             string   `json:"description"`
		FiToFiInformationLine1                  string   `json:"fi_to_fi_information_line_1"`
		FiToFiInformationLine2                  string   `json:"fi_to_fi_information_line_2"`
		FiToFiInformationLine3                  string   `json:"fi_to_fi_information_line_3"`
		FiToFiInformationLine4                  string   `json:"fi_to_fi_information_line_4"`
		FiToFiInformationLine5                  string   `json:"fi_to_fi_information_line_5"`
		FiToFiInformationLine6                  string   `json:"fi_to_fi_information_line_6"`
		ID                                      string   `json:"id"`
		IdempotencyKey                          string   `json:"idempotency_key"`
		Imad                                    string   `json:"imad"`
		InitiatedAt                             string   `json:"initiated_at"`
		IsIncoming                              bool     `json:"is_incoming"`
		IsOnUs                                  bool     `json:"is_on_us"`
		ManualReviewAt                          *string  `json:"manual_review_at"`
		Omad                                    string   `json:"omad"`
		OriginatorAccountNumber                 string   `json:"originator_account_number"`
		OriginatorName                          string   `json:"originator_name"`
		OriginatorToBeneficiaryInformationLine1 string   `json:"originator_to_beneficiary_information_line_1"`
		OriginatorToBeneficiaryInformationLine2 string   `json:"originator_to_beneficiary_information_line_2"`
		OriginatorToBeneficiaryInformationLine3 string   `json:"originator_to_beneficiary_information_line_3"`
		OriginatorToBeneficiaryInformationLine4 string   `json:"originator_to_beneficiary_information_line_4"`
		PendingSubmissionAt                     *string  `json:"pending_submission_at,omitempty"`
		PreviousMessageReference                string   `json:"previous_message_reference"`
		RawBeneficiaryAddress                   string   `json:"raw_beneficiary_address"`
		RawOriginatorAddress                    string   `json:"raw_originator_address"`
		ReceiverDiName                          string   `json:"receiver_di_name"`
		ReceiverDiRoutingNumber                 string   `json:"receiver_di_routing_number"`
		RejectedAt                              *string  `json:"rejected_at"`
		ReversalPairTransferID                  string   `json:"reversal_pair_transfer_id"`
		SenderAddress                           *Address `json:"sender_address,omitempty"`
		SenderDiName                            string   `json:"sender_di_name"`
		SenderDiRoutingNumber                   string   `json:"sender_di_routing_number"`
		SenderReference                         string   `json:"sender_reference"`
		Status                                  string   `json:"status"`
		SubmittedAt                             *string  `json:"submitted_at,omitempty"`
		SubtypeCode                             string   `json:"subtype_code"`
		TypeCode                                string   `json:"type_code"`
		UpdatedAt                               string   `json:"updated_at"`
		WireDrawdownRequestID                   string   `json:"wire_drawdown_request_id"`
	}

	// IntlWireTransfer schema. International wire transfers are used to send/receive money to/from outside of the United States via the Swift network
	IntlWireTransfer struct {
		AccountNumberID            string      `json:"account_number_id"`
		AllowOverdraft             bool        `json:"allow_overdraft"`
		Amount                     int         `json:"amount"`
		BankAccountID              string      `json:"bank_account_id"`
		BeneficiaryAccountNumber   string      `json:"beneficiary_account_number"`
		BeneficiaryAddress         *Address    `json:"beneficiary_address,omitempty"`
		BeneficiaryFi              string      `json:"beneficiary_fi"`
		BeneficiaryName            string      `json:"beneficiary_name"`
		CancellationReason         *string     `json:"cancellation_reason"`
		CancellationStatus         *string     `json:"cancellation_status"`
		ChargeBearer               string      `json:"charge_bearer"`
		Charges                    interface{} `json:"charges"`
		ColumnFixedFee             int         `json:"column_fixed_fee"`
		CompletedAt                string      `json:"completed_at"`
		CounterpartyID             string      `json:"counterparty_id"`
		CreatedAt                  string      `json:"created_at"`
		CurrencyCode               string      `json:"currency_code"`
		Description                string      `json:"description"`
		EndToEndID                 string      `json:"end_to_end_id"`
		FxQuoteID                  string      `json:"fx_quote_id"`
		FxRate                     interface{} `json:"fx_rate"`
		ID                         string      `json:"id"`
		IdempotencyKey             *string     `json:"idempotency_key"`
		InitiatedAt                string      `json:"initiated_at"`
		InstructedAmount           int         `json:"instructed_amount"`
		InstructedCurrencyCode     string      `json:"instructed_currency_code"`
		InstructionID              string      `json:"instruction_id"`
		InstructionToBeneficiaryFi string      `json:"instruction_to_beneficiary_fi"`
		IntermediaryFis            []string    `json:"intermediary_fis"`
		IsIncoming                 bool        `json:"is_incoming"`
		ManualReviewAt             *string     `json:"manual_review_at"`
		OriginatorAccountNumber    string      `json:"originator_account_number"`
		OriginatorAddress          *Address    `json:"originator_address,omitempty"`
		OriginatorFi               string      `json:"originator_fi"`
		OriginatorName             string      `json:"originator_name"`
		PendingSubmissionAt        string      `json:"pending_submission_at"`
		PlatformFixedFee           int         `json:"platform_fixed_fee"`
		PlatformFxFee              int         `json:"platform_fx_fee"`
		RawMessage                 string      `json:"raw_message"`
		RemittanceInfo             interface{} `json:"remittance_info"`
		ReturnReason               *string     `json:"return_reason"`
		ReturnedAmount             *int        `json:"returned_amount"`
		ReturnedAt                 *string     `json:"returned_at"`
		ReturnedCurrencyCode       *string     `json:"returned_currency_code"`
		SettledAmount              int         `json:"settled_amount"`
		SettledCurrencyCode        string      `json:"settled_currency_code"`
		SettlementDate             string      `json:"settlement_date"`
		Status                     string      `json:"status"`
		SubmittedAt                *string     `json:"submitted_at,omitempty"`
		Uetr                       string      `json:"uetr"`
		UltimateBeneficiaryAddress *string     `json:"ultimate_beneficiary_address"`
		UltimateBeneficiaryName    string      `json:"ultimate_beneficiary_name"`
		UltimateOriginatorAddress  *string     `json:"ultimate_originator_address"`
		UltimateOriginatorName     string      `json:"ultimate_originator_name"`
		UpdatedAt                  string      `json:"updated_at"`
	}

	// CreateCounterPartyRequest to create a new counterparty
	CreateCounterPartyRequest struct {
		Email              string  `json:"email,omitempty"`
		Name               string  `json:"name"`
		Phone              string  `json:"phone,omitempty"`
		RoutingNumber      string  `json:"routing_number"`
		RoutingNumberType  string  `json:"routing_number_type,omitempty"`
		AccountNumber      string  `json:"account_number"`
		AccountType        string  `json:"account_type,omitempty"`
		Address            Address `json:"address,omitempty"`
		Description        string  `json:"description"`
		LegalID            string  `json:"legal_id,omitempty"`
		LegalType          string  `json:"legal_type,omitempty"`
		LocalAccountNumber string  `json:"local_account_number,omitempty"`
		LocalBankCode      string  `json:"local_bank_code,omitempty"`
	}

	// CounterParty is the legal entity on the other side of the transaction
	CounterParty struct {
		ID                   string           `json:"id"`
		Name                 string           `json:"name"`
		Phone                string           `json:"phone"`
		Email                string           `json:"email"`
		Description          string           `json:"description"`
		AccountNumber        string           `json:"account_number"`
		AccountType          string           `json:"account_type"`
		Address              Address          `json:"address"`
		IsColumnAccount      bool             `json:"is_column_account"`
		LegalID              string           `json:"legal_id"`
		LegalType            string           `json:"legal_type"`
		LocalAccountNumber   string           `json:"local_account_number"`
		LocalBankCode        string           `json:"local_bank_code"`
		LocalBankCountryCode string           `json:"local_bank_country_code"`
		LocalBankName        string           `json:"local_bank_name"`
		RoutingNumber        string           `json:"routing_number"`
		RoutingNumberType    string           `json:"routing_number_type"`
		Wire                 CounterPartyWire `json:"wire"`
		WireDrawDownAllowed  bool             `json:"wire_drawdown_allowed"`
		CreatedAt            string           `json:"created_at"`
		UpdatedAt            string           `json:"updated_at"`
	}

	// CounterPartyWire schema
	CounterPartyWire struct {
		BeneficiaryAddress Address `json:"beneficiary_address"`
		BeneficiaryEmail   string  `json:"beneficiary_email"`
		BeneficiaryLegalID string  `json:"beneficiary_legal_id"`
		BeneficiaryName    string  `json:"beneficiary_name"`
		BeneficiaryPhone   string  `json:"beneficiary_phone"`
		BeneficiaryType    string  `json:"beneficiary_type"`
		LocalAccountNumber string  `json:"local_account_number"`
		LocalBankCode      string  `json:"local_bank_code"`
	}

	// Institution schema returned to get the information about a bank
	Institution struct {
		AchEligible        bool   `json:"ach_eligible"`
		City               string `json:"city"`
		CountryCode        string `json:"country_code"`
		CreatedAt          string `json:"created_at"`
		FullName           string `json:"full_name"`
		PhoneNumber        string `json:"phone_number"`
		RoutingNumber      string `json:"routing_number"`
		RoutingNumberType  string `json:"routing_number_type"`
		ShortName          string `json:"short_name"`
		State              string `json:"state"`
		StreetAddress      string `json:"street_address"`
		UpdatedAt          string `json:"updated_at"`
		WireEligible       bool   `json:"wire_eligible"`
		WireSettlementOnly bool   `json:"wire_settlement_only"`
		ZipCode            string `json:"zip_code"`
	}

	// BookTransfer is the current state of a single book transfer initiated in Column. A book transfer is the movement of funds between two bank accounts under your platform, and once initiated, happen instantaneously 24/7.
	BookTransfer struct {
		ID                      string `json:"id"`
		AllowOverdraft          bool   `json:"allow_overdraft"`
		Amount                  int    `json:"amount"`
		CurrencyCode            string `json:"currency_code"`
		Description             string `json:"description"`
		IdempotencyKey          string `json:"idempotency_key"`
		ReceiverAccountNumberID string `json:"receiver_account_number_id"`
		ReceiverBankAccountID   string `json:"receiver_bank_account_id"`
		SenderAccountNumberID   string `json:"sender_account_number_id"`
		SenderBankAccountID     string `json:"sender_bank_account_id"`
		Status                  string `json:"status"`
		CreatedAt               string `json:"created_at"`
		UpdatedAt               string `json:"updated_at"`
	}

	// CreateBookTransfer initiates a book transfer, a transaction between two Column accounts that you own.
	CreateBookTransfer struct {
		Amount            int64  `json:"amount"`
		Currency          string `json:"currency_code"`
		Description       string `json:"description"`
		SenderAccountID   string `json:"sender_bank_account_id"`
		ReceiverAccountID string `json:"receiver_bank_account_id"`
		Reference         string `json:"-"`
	}

	// FxQuote - The Foreign Exchange Quote object represents the rate quote for an international wire transfer.
	FxQuote struct {
		BuyAmount        int64  `json:"buy_amount"`
		BuyCurrencyCode  string `json:"buy_currency_code"`
		CreatedAt        string `json:"created_at"`
		ExpiredAt        string `json:"expired_at"`
		ID               string `json:"id"`
		Rate             string `json:"rate"`
		RateDate         string `json:"rate_date"`
		SellAmount       int64  `json:"sell_amount"`
		SellCurrencyCode string `json:"sell_currency_code"`
		Status           string `json:"status"`
		UpdatedAt        string `json:"updated_at"`
	}

	// CreateFxQuoteRequest schema to create or request for an FX quote
	CreateFxQuoteRequest struct {
		BuyAmount       int64  `json:"buy_amount"`
		BuyCurrencyCode string `json:"buy_currency_code"`
		RateDate        string `json:"rate_date,omitempty"`
		BookDirectly    *bool  `json:"book_directly,omitempty"`
	}

	// RejectReason is the reason for rejecting a transaction
	RejectReason string

	// CancelReason is the reason for canceling a transfer
	CancelReason string
)

// GetMt103 extracts the MT103 details
func (i *IntlWireTransfer) GetMt103() []string {
	var mt103Messages []string

	date := ""
	// extract date
	t, err := time.Parse(time.RFC3339, i.InitiatedAt)
	if err == nil {
		// Extract the day and month
		date = fmt.Sprintf("%02d%02d", t.Day(), t.Month())
	}

	mt103Messages = append(mt103Messages, fmt.Sprintf(":20:%s", i.EndToEndID))
	mt103Messages = append(mt103Messages, ":23B:CRED")
	mt103Messages = append(mt103Messages, fmt.Sprintf(":32A:%s%s%d,%02d", date, i.InstructedCurrencyCode, i.InstructedAmount/100, i.InstructedAmount%100))
	mt103Messages = append(mt103Messages, fmt.Sprintf(":33B:%s%d,%02d", i.InstructedCurrencyCode, i.InstructedAmount/100, i.InstructedAmount%100))
	mt103Messages = append(mt103Messages, fmt.Sprintf(":50K:%s", i.OriginatorAccountNumber))
	mt103Messages = append(mt103Messages, i.OriginatorName)
	mt103Messages = append(mt103Messages, fmt.Sprintf("%s, %s, %s, %s", i.OriginatorAddress.Line1, i.OriginatorAddress.City, i.OriginatorAddress.PostalCode, i.OriginatorAddress.CountryCode))
	if len(i.IntermediaryFis) > 0 {
		sendersCorrespondentBank := i.IntermediaryFis[0]
		if len(sendersCorrespondentBank) > 0 {
			mt103Messages = append(mt103Messages, fmt.Sprintf(":53A:%s", sendersCorrespondentBank))
		}

		if len(i.IntermediaryFis) > 1 {
			if len(i.IntermediaryFis[1]) > 0 {
				mt103Messages = append(mt103Messages, fmt.Sprintf(":56A:%s", i.IntermediaryFis[1]))
			}
		}
	}
	mt103Messages = append(mt103Messages, fmt.Sprintf(":57A:%s", i.BeneficiaryFi))
	mt103Messages = append(mt103Messages, fmt.Sprintf(":59:%s", i.BeneficiaryAccountNumber))
	mt103Messages = append(mt103Messages, i.BeneficiaryName)
	mt103Messages = append(mt103Messages, fmt.Sprintf("%s, %s %s", i.BeneficiaryAddress.Line1, i.BeneficiaryAddress.City, i.BeneficiaryAddress.PostalCode))
	mt103Messages = append(mt103Messages, fmt.Sprintf(":70:%s", i.Description))
	mt103Messages = append(mt103Messages, ":71A:SHA")

	return mt103Messages
}

const (
	// ReasonInvalidBeneficiaryAccNumber for when the beneficiary account number is invalid
	ReasonInvalidBeneficiaryAccNumber RejectReason = "invalid_beneficiary_account_number"

	// ReasonBeneficiaryMisMatch for when beneficiary information in the wire does not match the corresponding information on the account
	ReasonBeneficiaryMisMatch RejectReason = "beneficiary_mismatch"

	// ReasonIncorrectAmount for when transfer amount is not the same as expected by the beneficiary
	ReasonIncorrectAmount RejectReason = "incorrect_amount"

	// ReasonRefusedByBeneficiary for when the beneficiary refuses to accept this transfer
	ReasonRefusedByBeneficiary RejectReason = "refused_by_beneficiary"

	// ReasonCancellationRequested for when the cancellation request from the originator is approved
	ReasonCancellationRequested RejectReason = "cancellation_requested"

	// ReasonFraud for when the transfer is a fraud
	ReasonFraud RejectReason = "fraud"

	// ReasonComplianceRejected for when the transfer did not pass compliance check
	ReasonComplianceRejected RejectReason = "compliance_rejected"

	// CancelReasonIncorrectAmount cancel reason
	CancelReasonIncorrectAmount CancelReason = "incorrect_amount"
	// CancelReasonIncorrectCurrency cancel reason
	CancelReasonIncorrectCurrency CancelReason = "incorrect_currency"
	// CancelReasonDuplicate cancel reason
	CancelReasonDuplicate CancelReason = "duplicate"
	// CancelReasonTechFailure cancel reason
	CancelReasonTechFailure CancelReason = "tech_failure"
	// CancelReasonPaymentNotJustified cancel reason
	CancelReasonPaymentNotJustified CancelReason = "payment_not_justified"
	// CancelReasonRequestedByOriginator cancel reason
	CancelReasonRequestedByOriginator CancelReason = "requested_by_originator"

	// ColumnSwiftCode is the SWIFT code for COLUMN
	ColumnSwiftCode = "CLNOUS66"

	// RoutingNumberTypeABA is the string representation of routing number aba
	RoutingNumberTypeABA = "aba"
	// RoutingNumberTypeBIC is the string representation of routing number bic (swift)
	RoutingNumberTypeBIC = "bic"

	// WireTransferIDPrefix is the prefix of all IDs of wire transfers on the column platform
	WireTransferIDPrefix = "wire_"
	// IntlWireTransferIDPrefix is the prefix of all IDs of int'l wire transfers on the column platform
	IntlWireTransferIDPrefix = "swft_"
	// BookTransferIDPrefix is the prefix of all IDs of book transfers on the column platform
	BookTransferIDPrefix = "book_"
)
