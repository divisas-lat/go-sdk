package divisas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/divisas-lat/go-sdk/cache"
	"github.com/divisas-lat/go-sdk/models"
)

const (
	DefaultBaseURL = "https://api.divisas.lat/v1"
	DefaultCacheTTL = 3600 * time.Second
)

// Client represents the Divisas.lat API client
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	Cache      *cache.MemoryCache
	CacheTTL   time.Duration
}

// Option represents a functional option for configuring the Client
type Option func(*Client)

// WithAPIKey sets the API key explicitly
func WithAPIKey(key string) Option {
	return func(c *Client) {
		c.APIKey = key
	}
}

// WithBaseURL sets a custom base URL
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}

// WithCacheTTL sets the duration for which responses are cached
func WithCacheTTL(ttl time.Duration) Option {
	return func(c *Client) {
		c.CacheTTL = ttl
	}
}

// NewClient creates and returns a new Divisas.lat API client
func NewClient(opts ...Option) *Client {
	client := &Client{
		APIKey:     os.Getenv("DIVISAS_API_KEY"),
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Cache:      cache.NewMemoryCache(),
		CacheTTL:   DefaultCacheTTL,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// Query returns a new QueryBuilder
func (c *Client) Query() *QueryBuilder {
	return &QueryBuilder{
		client: c,
	}
}

// request executes the HTTP request, managing authentication and caching
func (c *Client) request(ctx context.Context, method, endpoint string, query url.Values, target interface{}) error {
	reqURL := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	if len(query) > 0 {
		reqURL = fmt.Sprintf("%s?%s", reqURL, query.Encode())
	}

	// Try cache for GET requests
	if method == http.MethodGet && c.CacheTTL > 0 {
		if cachedData := c.Cache.Get(reqURL); cachedData != nil {
			return json.Unmarshal(cachedData, target)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, reqURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errorResponse struct {
			Message string `json:"message"`
		}
		if err := json.Unmarshal(body, &errorResponse); err == nil && errorResponse.Message != "" {
			return &models.DivisasError{StatusCode: resp.StatusCode, Message: errorResponse.Message}
		}
		return &models.DivisasError{StatusCode: resp.StatusCode, Message: string(bytes.TrimSpace(body))}
	}

	if err := json.Unmarshal(body, target); err != nil {
		return err
	}

	// Save to cache on success
	if method == http.MethodGet && c.CacheTTL > 0 {
		c.Cache.Set(reqURL, body, c.CacheTTL)
	}

	return nil
}
