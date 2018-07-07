package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Client struct {
	ID                uint
	AccountName       string `json:"account_name"`
	AccountNumber     string `gorm:"unique" json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

type SenderCharge struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type FXType struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

type Sponsor struct {
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
}

type ChargesInformationType struct {
	BearerCode              string         `json:"bearer_code"`
	SenderCharges           []SenderCharge `json:"sender_charges"`
	ReceiverChargesAmount   string         `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string         `json:"receiver_charges_currency"`
}

type PaymentAttributes struct {
	Amount               string                 `json:"amount"`
	BeneficiaryParty     Client                 `json:"beneficiary_party"`
	ChargesInformation   ChargesInformationType `gorm:"embedded" json:"charges_information"`
	Currency             string                 `json:"currency"`
	DebtorParty          Client                 `json:"debtor_party"`
	EndToEndReference    string                 `json:"end_to_end_reference"`
	FX                   FXType                 `gorm:"embedded" json:"fx"`
	NumericReference     string                 `json:"numeric_reference"`
	PaymentID            string                 `json:"payment_id"`
	PaymentPurpose       string                 `json:"payment_purpose"`
	PaymentScheme        string                 `json:"payment_scheme"`
	PaymentType          string                 `json:"payment_type"`
	ProcessingDate       string                 `json:"processing_date"`
	Reference            string                 `json:"reference"`
	SchemaPaymentSubType string                 `json:"scheme_payment_sub_type"`
	SchemaPaymentType    string                 `json:"scheme_payment_type"`
}

type Payment struct {
	ID             uint
	MainPaymentID  string            `json:"id"`
	Type           string            `json:"type"`
	Version        int               `json:"version"`
	OrganisationID string            `json:"organisation_id"`
	Attributes     PaymentAttributes `gorm:"embedded" json:"attributes"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.LogMode(true)
	db.AutoMigrate(&Payment{})
	return db
}
