package divisas_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/divisas-lat/go-sdk"
	"github.com/divisas-lat/go-sdk/enums"
)

func TestClient_GetToday(t *testing.T) {
	mockResponse := `{
		"country": "GT",
		"base_currency": "GTQ",
		"date": "2026-05-05",
		"source": "Banguat",
		"cached": false,
		"rate": {
			"currency_code": "USD",
			"buy": 7.63,
			"sell": 7.63
		}
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/GT/rates" {
			t.Errorf("Expected path /v1/GT/rates, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("Expected Bearer test-key, got %s", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := divisas.NewClient(
		divisas.WithAPIKey("test-key"),
		divisas.WithBaseURL(server.URL+"/v1"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.Query().ForCountry(enums.Guatemala).GetToday(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Country != "GT" {
		t.Errorf("Expected country GT, got %s", res.Country)
	}
	if res.Rate == nil {
		t.Fatalf("Expected rate to not be nil")
	}
	if res.Rate.Buy != 7.63 {
		t.Errorf("Expected buy rate 7.63, got %f", res.Rate.Buy)
	}
}

func TestClient_Convert(t *testing.T) {
	mockResponse := `{
		"from": {
			"currency": "USD",
			"amount": 100
		},
		"to": {
			"currency": "GTQ",
			"amount": 763
		},
		"amount": 100,
		"result": 763,
		"effective_rate": 7.63,
		"via": "Banguat",
		"date": "2026-05-05",
		"note": ""
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/GT/rates/convert" {
			t.Errorf("Expected path /v1/GT/rates/convert, got %s", r.URL.Path)
		}
		if r.URL.Query().Get("currency") != "USD" {
			t.Errorf("Expected query currency=USD, got %s", r.URL.Query().Get("currency"))
		}
		if r.URL.Query().Get("to") != "GTQ" {
			t.Errorf("Expected query to=GTQ, got %s", r.URL.Query().Get("to"))
		}
		if r.URL.Query().Get("amount") != "100" {
			t.Errorf("Expected query amount=100, got %s", r.URL.Query().Get("amount"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := divisas.NewClient(
		divisas.WithBaseURL(server.URL+"/v1"),
	)

	ctx := context.Background()
	res, err := client.Query().
		ForCountry(enums.Guatemala).
		WithCurrency(enums.USD).
		Convert(ctx, enums.GTQ, 100)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Result != 763 {
		t.Errorf("Expected result 763, got %f", res.Result)
	}
}
