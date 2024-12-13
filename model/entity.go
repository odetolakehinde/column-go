package model

type (
	// CreatePersonRequest payload request for creating a new person request
	CreatePersonRequest struct {
		FirstName   string   `json:"first_name"`
		LastName    string   `json:"last_name"`
		MiddleName  string   `json:"middle_name"`
		SSN         string   `json:"ssn,omitempty"`
		Passport    Identity `json:"passport,omitempty"`
		DateOfBirth string   `json:"date_of_birth"`
		Email       string   `json:"email"`
		Address     Address  `json:"address"`
	}

	// PersonEntity schema represents a person entity
	PersonEntity struct {
		ID                 string          `json:"id"`
		IsRoot             bool            `json:"is_root"`
		Type               string          `json:"type"`
		PersonDetails      PersonalDetails `json:"person_details"`
		ReviewReasons      []interface{}   `json:"review_reasons"`
		VerificationStatus string          `json:"verification_status"`
		Documents          []interface{}   `json:"documents"`
	}

	// Address schema
	Address struct {
		City        string `json:"city"`
		CountryCode string `json:"country_code"`
		Line1       string `json:"line_1"`
		Line2       string `json:"line_2"`
		PostalCode  string `json:"postal_code"`
		State       string `json:"state"`
	}

	// PersonalDetails schema
	PersonalDetails struct {
		Address        Address  `json:"address"`
		DateOfBirth    string   `json:"date_of_birth"`
		Email          string   `json:"email"`
		FirstName      string   `json:"first_name"`
		LastName       string   `json:"last_name"`
		MiddleName     string   `json:"middle_name"`
		SSN            string   `json:"ssn,omitempty"`
		Passport       Identity `json:"passport,omitempty"`
		DriversLicense Identity `json:"drivers_license,omitempty"`
		NationalID     Identity `json:"national_id,omitempty"`
	}

	// BusinessEntity schema
	BusinessEntity struct {
		BusinessDetails struct {
			Address          Address           `json:"address"`
			BeneficialOwners []PersonalDetails `json:"beneficial_owners"`
			BusinessName     string            `json:"business_name"`
			EIN              string            `json:"ein,omitempty"`
			RegistrationID   Identity          `json:"registration_id,omitempty"`
			Industry         string            `json:"industry"`
			LegalType        string            `json:"legal_type"`
			Website          string            `json:"website"`
		} `json:"business_details"`
		Documents          []interface{} `json:"documents"`
		ID                 string        `json:"id"`
		IsRoot             bool          `json:"is_root"`
		ReviewReasons      []interface{} `json:"review_reasons"`
		Type               string        `json:"type"`
		VerificationStatus string        `json:"verification_status"`
	}

	// CreateBusinessRequest schema
	CreateBusinessRequest struct {
		EIN              string                   `json:"ein,omitempty"`
		RegistrationID   Identity                 `json:"registration_id,omitempty"`
		BusinessName     string                   `json:"business_name"`
		Website          string                   `json:"website"`
		LegalType        string                   `json:"legal_type"`
		Address          Address                  `json:"address"`
		BeneficialOwners []BeneficialOwnerRequest `json:"beneficial_owners"`
	}

	// BeneficialOwnerRequest schema
	BeneficialOwnerRequest struct {
		DateOfBirth       string    `json:"date_of_birth"`
		FirstName         string    `json:"first_name"`
		LastName          string    `json:"last_name"`
		SSN               string    `json:"ssn,omitempty"`
		Passport          *Identity `json:"passport,omitempty"`
		DriversLicense    *Identity `json:"drivers_license,omitempty"`
		NationalID        *Identity `json:"national_id,omitempty"`
		Email             string    `json:"email"`
		IsControlPerson   bool      `json:"is_control_person"`
		IsBeneficialOwner bool      `json:"is_beneficial_owner"`
		JobTitle          string    `json:"job_title"`
		Address           Address   `json:"address"`
	}

	// Identity schema for identity details - drivers' license, ID etc
	Identity struct {
		Number      string `json:"number,omitempty"`
		CountryCode string `json:"country_code,omitempty"`
	}
)

// ErrorDictionary maps error codes to custom error descriptions.
var ErrorDictionary = map[string]string{
	"invalid_string_character":    string(ErrorCodeInvalidStringCharacter),
	"mandatory_parameter_missing": string(ErrorCodeMandatoryParameterMissing),
	"invalid_resource_id_format":  string(ErrorCodeInvalidResourceIDFormat),
	// Add more error codes and messages as needed.
}

const (
	// LegalTypeLimitedPartnership string representation of limited-partnership
	LegalTypeLimitedPartnership = "limited-partnership"

	// LegalTypeTrust string representation of trust
	LegalTypeTrust = "trust"

	// LegalTypeSoleProprietorship string representation of sole-proprietorship
	LegalTypeSoleProprietorship = "sole-proprietorship"

	// LegalTypeCorporation string representation of corporation
	LegalTypeCorporation = "corporation"

	// LegalTypeLLC string representation of llc
	LegalTypeLLC = "llc"

	// LegalTypeGeneralPartnership string representation of general-partnership
	LegalTypeGeneralPartnership = "general-partnership"

	// LegalTypeProfessionalAssociation string representation of professional-association
	LegalTypeProfessionalAssociation = "professional-association"

	// LegalTypeOthers string representation of others
	LegalTypeOthers = "others"

	// ErrorCodeInvalidStringCharacter string representation of invalid string character on Column
	ErrorCodeInvalidStringCharacter ColumnErrorCode = "Invalid string character"

	// ErrorCodeInvalidResourceIDFormat string representation of  Invalid Resource Id Format
	ErrorCodeInvalidResourceIDFormat ColumnErrorCode = "Invalid resource Id format"

	// ErrorCodeMandatoryParameterMissing string representation of mandatory parameter missing
	ErrorCodeMandatoryParameterMissing ColumnErrorCode = "Mandatory parameter missing"
)
