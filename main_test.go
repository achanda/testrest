package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"

	"github.com/achanda/testrest/app"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestEmptyPayment(t *testing.T) {
	mocket.Catcher.Register()

	// An empty list of payments should return []
	db, _ := gorm.Open(mocket.DRIVER_NAME, "any_string")
	router := mux.NewRouter()
	app := app.App{router, db}
	req, err := http.NewRequest("GET", "http://localhost/payments", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()
	app.GetAllPayments(w, req)
	assert.Equal(t, 200, w.Code, "OK response is expected")
	got := string(w.Body.Bytes())
	assert.Equal(t, 2, len(got))
	assert.JSONEq(t, "[]", got)
}

func TestCreateGetDeletePayment(t *testing.T) {
	data := []byte(`
	{"type":"Payment","id":"4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43","version":0,
	"organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb","attributes":
	{"amount":"100.21","beneficiary_party":{"account_name":"W Owens",
	"account_number":"31926819","account_number_code":"BBAN","account_type":0,
	"address":"1 The Beneficiary Localtown SE2","bank_id":"403000","bank_id_code":
	"GBDSC","name":"Wilfred Jeremiah Owens"},"charges_information":{"bearer_code":
	"SHAR","sender_charges":[{"amount":"5.00","currency":"GBP"},{"amount":"10.00",
	"currency":"USD"}],"receiver_charges_amount":"1.00","receiver_charges_currency":"USD"},
	"currency":"GBP","debtor_party":{"account_name":"EJ Brown Black","account_number":
	"GB29XABC10161234567801","account_number_code":"IBAN","address":"10 Debtor Crescent Sourcetown NE1",
	"bank_id":"203301","bank_id_code":"GBDSC","name":"Emelia Jane Brown"},"end_to_end_reference":
	"Wil piano Jan","fx":{"contract_reference":"FX123","exchange_rate":"2.00000","original_amount":
	"200.42","original_currency":"USD"},"numeric_reference":"1002001","payment_id":"123456789012345678",
	"payment_purpose":"Paying for goods/services","payment_scheme":"FPS","payment_type":"Credit",
	"processing_date":"2017-01-18","reference":"Payment for Ems piano lessons","scheme_payment_sub_type":
	"InternetBanking","scheme_payment_type":"ImmediatePayment","sponsor_party":{"account_number":"56781234",
	"bank_id":"123123","bank_id_code":"GBDSC"}}}
	`)
	mocket.Catcher.Register()
	db, _ := gorm.Open(mocket.DRIVER_NAME, "any_string")
	router := mux.NewRouter()
	app := app.App{router, db}

	// Test creating a new payment
	req, err := http.NewRequest("POST", "http://localhost/payments", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Test failed")
	}
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	app.CreatePayment(w, req)
	assert.Equal(t, http.StatusCreated, w.Code, "Expected 201")
}
