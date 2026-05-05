package divisas

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/divisas-lat/go-sdk/enums"
	"github.com/divisas-lat/go-sdk/models"
)

// QueryBuilder provides a fluent interface for constructing API requests
type QueryBuilder struct {
	client   *Client
	country  enums.Country
	currency enums.Currency
	source   string
}

// ForCountry sets the country for the query
func (qb *QueryBuilder) ForCountry(country enums.Country) *QueryBuilder {
	qb.country = country
	return qb
}

// WithCurrency sets an optional target currency for the query
func (qb *QueryBuilder) WithCurrency(currency enums.Currency) *QueryBuilder {
	qb.currency = currency
	return qb
}

// WithSource sets an optional source provider
func (qb *QueryBuilder) WithSource(source string) *QueryBuilder {
	qb.source = source
	return qb
}

func (qb *QueryBuilder) buildQuery() url.Values {
	q := url.Values{}
	if qb.currency != "" {
		q.Set("currency", string(qb.currency))
	}
	if qb.source != "" {
		q.Set("source", qb.source)
	}
	return q
}

// GetToday retrieves the current exchange rates
func (qb *QueryBuilder) GetToday(ctx context.Context) (*models.TodayRatesResponse, error) {
	if qb.country == "" {
		return nil, errors.New("country is required, use ForCountry()")
	}

	endpoint := fmt.Sprintf("/%s/rates", qb.country)
	if qb.currency != "" {
		endpoint = fmt.Sprintf("/%s/rates/%s", qb.country, qb.currency)
	}
	q := qb.buildQuery()
	q.Del("currency") // already in path

	var response models.TodayRatesResponse
	err := qb.client.request(ctx, http.MethodGet, endpoint, q, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Convert performs a currency conversion
func (qb *QueryBuilder) Convert(ctx context.Context, targetCurrency enums.Currency, amount float64) (*models.ConversionResponse, error) {
	if qb.country == "" {
		return nil, errors.New("country is required, use ForCountry()")
	}
	if targetCurrency == "" {
		return nil, errors.New("target currency is required")
	}

	endpoint := fmt.Sprintf("/%s/rates/convert", qb.country)
	q := qb.buildQuery()
	q.Set("to", string(targetCurrency))
	q.Set("amount", strconv.FormatFloat(amount, 'f', -1, 64))

	var response models.ConversionResponse
	err := qb.client.request(ctx, http.MethodGet, endpoint, q, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetHistory retrieves historical exchange rates
func (qb *QueryBuilder) GetHistory(ctx context.Context, startDate, endDate string) (*models.HistoricalRateResponse, error) {
	if qb.country == "" {
		return nil, errors.New("country is required, use ForCountry()")
	}
	if qb.currency == "" {
		return nil, errors.New("currency is required for historical data, use WithCurrency()")
	}

	endpoint := fmt.Sprintf("/%s/rates/history", qb.country)
	q := qb.buildQuery()
	if startDate != "" {
		q.Set("from", startDate)
	}
	if endDate != "" {
		q.Set("to", endDate)
	}

	var response models.HistoricalRateResponse
	err := qb.client.request(ctx, http.MethodGet, endpoint, q, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetStats retrieves descriptive statistics for the exchange rates
func (qb *QueryBuilder) GetStats(ctx context.Context, period string) (*models.StatsResponse, error) {
	if qb.country == "" {
		return nil, errors.New("country is required, use ForCountry()")
	}

	endpoint := fmt.Sprintf("/%s/rates/stats", qb.country)
	q := qb.buildQuery()
	if period != "" {
		q.Set("period", period)
	}

	var response models.StatsResponse
	err := qb.client.request(ctx, http.MethodGet, endpoint, q, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetForecast predicts future exchange rates
func (qb *QueryBuilder) GetForecast(ctx context.Context, days int) (*models.ForecastResponse, error) {
	if qb.country == "" {
		return nil, errors.New("country is required, use ForCountry()")
	}

	endpoint := fmt.Sprintf("/%s/rates/forecast", qb.country)
	q := qb.buildQuery()
	if days > 0 {
		q.Set("days", strconv.Itoa(days))
	}

	var response models.ForecastResponse
	err := qb.client.request(ctx, http.MethodGet, endpoint, q, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPercentile analyzes where the current rate sits historically
func (qb *QueryBuilder) GetPercentile(ctx context.Context, period string) (*models.PercentileResponse, error) {
	if qb.country == "" {
		return nil, errors.New("country is required, use ForCountry()")
	}

	endpoint := fmt.Sprintf("/%s/rates/percentile", qb.country)
	q := qb.buildQuery()
	if period != "" {
		q.Set("period", period)
	}

	var response models.PercentileResponse
	err := qb.client.request(ctx, http.MethodGet, endpoint, q, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
