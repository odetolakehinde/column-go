// Package main execution point for this integration
package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	apiCalls "github.com/ovalfi/integrations-playground/column/api"
	"github.com/ovalfi/integrations-playground/column/model"
)

func main() {
	fmt.Println("Running Column SDK")
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("app starting")
	defer logger.Info().Msg("stopped")
	api := apiCalls.New(&logger, model.BaseURL, model.APIKey, model.ProxyIPAddressAndPort)
	api.RunInSandboxMode() // to ensure it is running in sandbox mode

	//ctx := context.Background()

	// CreatePersonEntity
	//personEntity, err := api.CreatePersonEntity(ctx, model.CreatePersonRequest{
	//	FirstName:   "Victor",
	//	LastName:    "Raji",
	//	MiddleName:  "Damisa",
	//	SSN:         "999999999",
	//	DateOfBirth: "1996-08-04",
	//	Email:       "victor@ovalfi.com",
	//	Address: model.Address{
	//		City:        "Lekki",
	//		CountryCode: "NG",
	//		Line1:       "Lekki Funds",
	//		Line2:       "",
	//		PostalCode:  "200289",
	//		State:       "Lagos",
	//	},
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.CreatePersonEntity failed.")
	//	return
	//}
	//logger.Info().Msgf("api.CreatePersonEntity result: %+v\n", personEntity)

	//personEntity, err := api.CreateBusinessEntity(context.Background(), model.CreateBusinessRequest{
	//	BusinessName: "Segsalerty Resources",
	//	Website:      "segsalerty.com",
	//	LegalType:    model.LegalTypeLimitedPartnership,
	//	Address: model.Address{
	//		Line1:       "Tinubu Street",
	//		City:        "Ikeja",
	//		State:       "Lagos",
	//		CountryCode: "NG",
	//	},
	//	RegistrationID: model.Identity{
	//		Number:      "1234566",
	//		CountryCode: "NG",
	//	},
	//	BeneficialOwners: []model.BeneficialOwnerRequest{
	//		{
	//			DateOfBirth: "1996-08-04",
	//			FirstName:   "Kehinde",
	//			LastName:    "ODETOLA",
	//			SSN:         "9999999-99",
	//			Passport: &model.Identity{
	//				Number:      "B12443",
	//				CountryCode: "NG",
	//			},
	//			Email:             "kehinde@ovalfi.com",
	//			IsControlPerson:   true,
	//			IsBeneficialOwner: true,
	//			JobTitle:          "Owner",
	//			Address: model.Address{
	//				Line1:       "Karpos Building, Ibadan",
	//				City:        "Ibadan",
	//				State:       "OYO",
	//				CountryCode: "NG",
	//				PostalCode:  "200283",
	//			},
	//		},
	//	},

	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.CreateBusinessEntity failed.")
	//	return
	//}
	//logger.Info().Msgf("api.CreateBusinessEntity result: %+v\n", personEntity)

	//personEntity, err := api.GetPersonEntity(context.Background(), "enti_2Uk7ItpueWAFeJbSZb6FSpHMrkQ") //
	//if err != nil {
	//	logger.Err(err).Msg("api.GetPersonEntity failed.")
	//	return
	//}
	//logger.Info().Msgf("api.GetPersonEntity result: %+v\n", personEntity)

	// CreateBank
	//bank, err := api.CreateBankAccount(context.Background(), model.CreateBankRequest{
	//	EntityID:    "enti_2UfFj5CUuryK0SzC1M16SS4a0rv",
	//	Description: uuid.New().String(),
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.CreateBank failed.")
	//	return
	//}
	//logger.Info().Msgf("api.CreateBank result: %+v\n", bank) // bacc_2UkKwqnZsGRKHVgwzUAWDh3fjSQ

	//bank, err := api.GetBankAccountByID(context.Background(), "acno_2c80U4Zmtt5vRMdeQGq4Ngyd5zm")
	//if err != nil {
	//	logger.Err(err).Msg("api.GetBankByID failed.")
	//	return
	//}
	//logger.Info().Msgf("api.GetBankByID result: %+v\n", bank)

	//bank, err := api.CreateBankAccountNumber(context.Background(), model.CreateBankAccountNoRequest{
	//	BankID:      "bacc_2UkKwqnZsGRKHVgwzUAWDh3fjSQ",
	//	Description: "some-random-account-number generated",
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.GetBankByID failed.")
	//	return
	//}
	//logger.Info().Msgf("api.GetBankByID result: %+v\n", bank)

	//bank, err := api.ListBankAccountNumbers(context.Background(), "bacc_2UkKwqnZsGRKHVgwzUAWDh3fjSQ")
	//if err != nil {
	//	logger.Err(err).Msg("api.ListBankAccountNumbers failed.")
	//	return
	//}
	//logger.Info().Msgf("api.ListBankAccountNumbers result: %+v\n", bank)

	//institution, err := api.GetInstitution(context.Background(), "HBUKGB4B") // 121145307
	//if err != nil {
	//	logger.Err(err).Msg("api.GetInstitution failed.")
	//	return
	//}
	//
	//logger.Info().Msgf("api.GetInstitution result: %+v\n", institution)

	//intlWire, err := api.InitiateIntlWireTransfer(context.Background(), model.CreateWireTransferRequest{
	//	Amount:           2324000,
	//	CurrencyCode:     "USD",
	//	Description:      "random stuff",
	//	CounterpartyID:   "cpty_2ajpdZslWgORMX3uMXucuQsy0uC",
	//	BankAccountID:    "bacc_2W4AcSfBfvUj9HHntejLcdBg9Ti",
	//	IntermediaryBank: "UNAFUS33XXX",
	//	RemittanceInfo: &model.RemittanceInfo{
	//		GeneralInfo:          "",
	//		BeneficiaryReference: "random stuff invoice",
	//	},
	//	Reference: "oriyomi-fx",
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.InitiateIntlWireTransfer failed.")
	//	return
	//}
	//
	//logger.Info().Msgf("api.InitiateIntlWireTransfer result: %+v\n", intlWire.GetMt103())

	//intlWire, err := api.GetIntlWireTransfer(context.Background(), "swft_2b7XC4nbESwgxGliW3T9tOb6dYq") // 121145307
	//if err != nil {
	//	logger.Err(err).Msg("api.GetIntlWireTransfer failed.")
	//	return
	//}
	//
	//logger.Info().Msgf("api.GetIntlWireTransfer result: %+v\n", intlWire.GetMt103())

	//intlWire, err := api.CancelOutgoingIntlWireTransfer(context.Background(), model.CancelIntlWireTransferRequest{
	//	ID:           "swft_2b7XC4nbESwgxGliW3T9tOb6dYq",
	//	CancelReason: model.CancelReasonPaymentNotJustified,
	//}) // 121145307
	//if err != nil {
	//	logger.Err(err).Msg("api.CancelOutgoingIntlWireTransfer failed.")
	//	return
	//}
	//
	//logger.Info().Msgf("api.CancelOutgoingIntlWireTransfer result: %+v\n", intlWire)

	//wire, err := api.GetWireTransferByID(context.Background(), "wire_2b4S7uR02NVLnmmSPx6h4JCMBPg") // 121145307
	//if err != nil {
	//	logger.Err(err).Msg("api.GetWireTransferByID failed.")
	//	return
	//}
	//
	//logger.Info().Msgf("api.GetWireTransferByID result: %+v\n", wire)

	//wire, err := api.GetAllWireTransfer(context.Background()) // 121145307
	//if err != nil {
	//	logger.Err(err).Msg("api.GetAllWireTransfer failed.")
	//	return
	//}
	//
	//logger.Info().Msgf("api.GetAllWireTransfer result: %+v\n", wire)

	//iWire, err := api.GetAllIntlWireTransfer(context.Background()) // 121145307
	//if err != nil {
	//	logger.Err(err).Msg("api.GetAllIntlWireTransfer failed.")
	//	return
	//}
	//
	//logger.Info().Msgf("api.GetAllIntlWireTransfer result: %+v\n", iWire)

	//intlWire, err := api.ReverseIncomingIntlWireTransfer(context.Background(), model.ReverseIntlWireTransferRequest{
	//	ID:           "swft_2VkOHCBhwsqUmQyOoJTsvpkZkIW",
	//	ReturnReason: model.ReasonBeneficiaryMisMatch,
	//	//MsgToOriginatorBank: "",
	//}) //
	//if err != nil {
	//	logger.Err(err).Msg("api.ReverseIncomingIntlWireTransfer failed.")
	//	return
	//}
	//logger.Info().Msgf("api.ReverseIncomingIntlWireTransfer result: %+v\n", intlWire)

	//err := api.SimulateWireInflow(context.Background(), model.SimulateWireTransferRequest{
	//	Amount:               100000,
	//	CurrencyCode:         "USD",
	//	DestinationAccNumber: "acno_2VXsSn3toZBAUFnjhBdDzZ90LC9",
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.SimulateWireInflow failed.")
	//	return
	//}

	//err := api.SimulateIntlWireInflow(context.Background(), model.SimulateWireTransferRequest{
	//	Amount:               100000,
	//	CurrencyCode:         "USD",
	//	DestinationAccNumber: "acno_2VXsSn3toZBAUFnjhBdDzZ90LC9",
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.SimulateWireInflow failed.")
	//	return
	//}

	//counterparty, err := api.CreateCounterparty(context.Background(), model.CreateCounterPartyRequest{
	//	//Email:              "",
	//	Name: "Odetola Kehinde", // "ORIYOMI FX",
	//	//Phone:              "",
	//	RoutingNumber: "104014138",
	//	//RoutingNumberType: "aba",
	//	AccountNumber: "11395454545",
	//	//AccountType:   "checking",
	//	Address: model.Address{
	//		City:        "AG",
	//		CountryCode: "NG",
	//		Line1:       "Southampton",
	//		Line2:       "",
	//		//PostalCode:  "2133425",
	//		//State:       "Oyo",
	//	},
	//	Description:        "2cecfae7-f618-4a42-a805-0050719c6645",
	//	LegalID:            "",
	//	LegalType:          "",
	//	LocalAccountNumber: "",
	//	LocalBankCode:      "",
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.CreateCounterparty failed.")
	//	return
	//}
	//logger.Info().Msgf("api.CreateCounterparty result: %+v\n", counterparty)

	// bank, err := api.InitiateBookTransfer(context.Background(), model.CreateBookTransfer{
	// 	Amount:            5000,
	// 	Currency:          "USD",
	// 	Description:       "some-random-description-here",
	// 	SenderAccountID:   "bacc_2VXsSoJ9eAHzrZRqBBX8cWqCPos",
	// 	ReceiverAccountID: "bacc_2VXv1ENbaY0RAAVeOcv1qIrFvoS",
	// })
	// if err != nil {
	// 	logger.Err(err).Msg("api.GetBankByID failed.")
	// 	return
	// }
	// logger.Info().Msgf("api.GetBankByID result: %+v\n", bank)

	// RUN FROM HERE FOR THE DB EXEC
	//ctx := context.Background()
	//
	//entities, err := api.GetBusinessEntities(ctx)
	//if err != nil {
	//	logger.Err(err).Msg("api.GetBusinessEntities failed.")
	//	return
	//}
	//
	//bankAccounts, err := api.GetBankAccounts(ctx)
	//if err != nil {
	//	logger.Err(err).Msg("api.GetBankAccounts failed.")
	//	return
	//}
	//logger.Info().Msgf("api.GetBankAccounts result: %+v\n", bankAccounts)
	//
	//// Convert the struct to JSON
	//entityData, err := json.MarshalIndent(entities, "", "    ")
	//if err != nil {
	//	fmt.Println("Error marshaling to JSON:", err)
	//	return
	//}
	//
	//// Convert the struct to JSON
	//bankAccountsData, err := json.MarshalIndent(bankAccounts, "", "    ")
	//if err != nil {
	//	fmt.Println("Error marshaling to JSON:", err)
	//	return
	//}
	//
	//// Save the JSON data to a file
	//err = ioutil.WriteFile("entities.json", entityData, 0644)
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//	return
	//}
	//
	//err = ioutil.WriteFile("bank-accounts.json", bankAccountsData, 0644)
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//	return
	//}

	//counterparties, err := api.ListCounterparties(ctx)
	//if err != nil {
	//	logger.Err(err).Msg("api.ListCounterparties failed.")
	//	return
	//}
	//
	//var scriptLines []string
	//
	//for _, data := range counterparties {
	//
	//	// get bank information
	//	institution, err := api.GetInstitution(ctx, data.RoutingNumber)
	//	if err != nil {
	//		fmt.Println("GetInstitution failed:", err)
	//		return
	//	}
	//
	//	var swiftCode, routingNo string
	//	if institution.RoutingNumberType == model.RoutingNumberTypeBIC {
	//		swiftCode = institution.RoutingNumber
	//	} else {
	//		routingNo = institution.RoutingNumber
	//	}
	//
	//	isWithinUS := "false"
	//	if institution.CountryCode == "US" {
	//		isWithinUS = "true"
	//	}
	//
	//	bankDetails := map[string]interface{}{
	//		"city":          data.Address.City,
	//		"country":       data.Address.CountryCode,
	//		"bankCode":      "",
	//		"bankName":      institution.FullName,
	//		"district":      institution.State,
	//		"swiftCode":     swiftCode,
	//		"bankBranch":    "",
	//		"isWithinUS":    isWithinUS,
	//		"postalCode":    institution.ZipCode,
	//		"accountName":   data.Name,
	//		"bankAddress":   fmt.Sprintf("%s, %s", institution.StreetAddress, institution.City),
	//		"accountNumber": data.AccountNumber,
	//		"isMobileMoney": "",
	//		"routingNumber": routingNo,
	//	}
	//
	//	personalDetails := map[string]interface{}{
	//		"city":       data.Address.City,
	//		"name":       data.Name,
	//		"address":    fmt.Sprintf("%s %s, %s", data.Address.Line1, data.Address.Line2, data.Address.City),
	//		"country":    data.Address.CountryCode,
	//		"district":   "",
	//		"postalCode": "",
	//	}
	//
	//	bankDetailsJSON, _ := json.Marshal(bankDetails)
	//	personalDetailsJSON, _ := json.Marshal(personalDetails)
	//
	//	scriptLine := fmt.Sprintf(`INSERT INTO "prod_banking".public.external_accounts (id, details, compliance_status, compliance_partner, status, bypass_reason, created_at, deactivated_at, updated_at, deleted_at, type, funds_transfer_ref, currency) VALUES (DEFAULT, '{"bankDetails": %s, "walletDetails": null, "personalDetails": %s, "intermediaryBank": null, "fundsTransferMethod": null}', 'no_action', 'oval', 'active', null, '%s', null, '%s', null, 'fiat_transfer', '{"column": "%s"}', 'USD');`,
	//		string(bankDetailsJSON), string(personalDetailsJSON), "now()", "now()", data.ID)
	//
	//	scriptLines = append(scriptLines, scriptLine)
	//}
	//
	//scriptContent := strings.Join(scriptLines, "\n")
	//err = ioutil.WriteFile("counterparties.sql", []byte(scriptContent), 0644)
	//if err != nil {
	//	fmt.Println("Error writing to file:", err)
	//	return
	//}

	//amount := 100000
	//currency := "EUR"
	//
	//_, err := api.CreateFxQuote(context.Background(), model.CreateFxQuoteRequest{
	//	BuyAmount:       int64(amount),
	//	BuyCurrencyCode: currency,
	//	//RateDate:        "12/12/2024",
	//	//BookDirectly:    nil,
	//})
	//if err != nil {
	//	logger.Err(err).Msg("api.CreateFxQuote failed.")
	//	return
	//}

	//fmt.Println("this is the quote --> ", quote)
	//
	//_, err = api.BookFxQuote(context.Background(), quote.ID)
	//if err != nil {
	//	logger.Err(err).Msg("api.BookFxQuote failed.")
	//	return
	//}

	//_, err = api.CancelFxQuote(context.Background(), quote.ID)
	//if err != nil {
	//	logger.Err(err).Msg("api.CancelFxQuote failed.")
	//	return
	//}

	//_, err = api.InitiateIntlWireTransfer(context.Background(), model.CreateWireTransferRequest{
	//	Amount:               int64(amount),
	//	CurrencyCode:         currency,
	//	Description:          "ninja-doings",
	//	CounterpartyID:       "cpty_2ajpdZslWgORMX3uMXucuQsy0uC",
	//	BankAccountID:        "bacc_2W4AcSfBfvUj9HHntejLcdBg9Ti",
	//	MsgToBeneficiaryBank: "ninja-is-at-it-again.mafo",
	//	IntermediaryBank:     "",
	//	RemittanceInfo: &model.RemittanceInfo{
	//		GeneralInfo:          "ninja-doings",
	//		BeneficiaryReference: "ninja-doings",
	//	},
	//	FxQuoteID: quote.ID,
	//})

	//_, err = api.GetFxQuote(context.Background(), quote.ID)
	//if err != nil {
	//	logger.Err(err).Msg("api.GetFxQuote failed.")
	//	return
	//}
}
