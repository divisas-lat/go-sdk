package models

// CurrencyRate represents the details of a single currency's exchange rate
type CurrencyRate struct {
	CurrencyCode string  `json:"currency_code"`
	Buy          float64 `json:"buy"`
	Sell         float64 `json:"sell"`
	CurrencyName *string `json:"currency_name,omitempty"`
	Date         *string `json:"date,omitempty"`
}

// TodayRatesResponse represents the response for current day's exchange rates
type TodayRatesResponse struct {
	Country       string         `json:"country"`
	BaseCurrency  string         `json:"base_currency"`
	Date          string         `json:"date"`
	Source        string         `json:"source"`
	Cached        bool           `json:"cached"`
	Rates         []CurrencyRate `json:"rates,omitempty"`
	Rate          *CurrencyRate  `json:"rate,omitempty"`
}

// HistoricalRateResponse represents historical rates data
type HistoricalRateResponse struct {
	Country      string         `json:"country"`
	CurrencyCode string         `json:"currency_code"`
	Source       string         `json:"source"`
	Total        int            `json:"total"`
	History      []CurrencyRate `json:"history"`
}

// ConversionResponse represents the result of a currency conversion
type ConversionResponse struct {
	From struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"from"`
	To struct {
		Currency string  `json:"currency"`
		Amount   float64 `json:"amount"`
	} `json:"to"`
	Amount        float64 `json:"amount"`
	Result        float64 `json:"result"`
	EffectiveRate float64 `json:"effective_rate"`
	Via           string  `json:"via"`
	Date          string  `json:"date"`
	Note          string  `json:"note"`
}

// StatsResponse represents statistical information about a currency
type StatsResponse struct {
	Country      string `json:"country"`
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
	Period       string `json:"period"`
	Current      struct {
		Buy  float64 `json:"buy"`
		Sell float64 `json:"sell"`
		Date string  `json:"date"`
	} `json:"current"`
	Stats struct {
		MinBuy     float64 `json:"min_buy"`
		MaxBuy     float64 `json:"max_buy"`
		AvgBuy     float64 `json:"avg_buy"`
		MinSell    float64 `json:"min_sell"`
		MaxSell    float64 `json:"max_sell"`
		AvgSell    float64 `json:"avg_sell"`
		Volatility float64 `json:"volatility"`
	} `json:"stats"`
	DataPoints int `json:"data_points"`
}

// ForecastResponse represents future predicted exchange rates
type ForecastResponse struct {
	Country      string `json:"country"`
	CurrencyCode string `json:"currency_code"`
	Model        string `json:"model"`
	BasedOnDays  int    `json:"based_on_days"`
	Forecast     []struct {
		Date string  `json:"date"`
		Buy  float64 `json:"buy"`
		Sell float64 `json:"sell"`
	} `json:"forecast"`
}

// PercentileResponse represents an analysis of where current rates fall within historical data
type PercentileResponse struct {
	Country      string `json:"country"`
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
	Today        struct {
		Date string  `json:"date"`
		Buy  float64 `json:"buy"`
		Sell float64 `json:"sell"`
	} `json:"today"`
	Percentile float64 `json:"percentile"`
	Period     string  `json:"period"`
	Signal     string  `json:"signal"`
	Range      struct {
		MinBuy  float64 `json:"min_buy"`
		MaxBuy  float64 `json:"max_buy"`
		MinSell float64 `json:"min_sell"`
		MaxSell float64 `json:"max_sell"`
	} `json:"range"`
}
