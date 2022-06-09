package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gerbenjacobs/go-habbo/habbo"
)

const version = "0.1.0"

// Parser is a Habbo API parser
type Parser struct {
	httpClient *http.Client
	userAgent  string
}

// NewParser creates a new Habbo API parser
func NewParser(httpClient *http.Client) *Parser {
	return &Parser{httpClient: httpClient, userAgent: "github.com/gerbenjacobs/go-habbo v" + version}
}

// WithUserAgent sets the user agent for the parser
func (p *Parser) WithUserAgent(userAgent string) {
	p.userAgent = userAgent
}

func (p *Parser) habbo(ctx context.Context, hotel string, identifier string, useUniqueID bool) (*habbo.Habbo, error) {
	// create url - name lookup vs habbo unique ID hhus-9cd61b156972c2eb33a145d69918f965
	var url = fmt.Sprintf("https://www.habbo.%s/api/public/users?name=%s", hotel, identifier)
	if useUniqueID {
		url = fmt.Sprintf("https://www.habbo.%s/api/public/users/%s", hotel, identifier)
	}

	// call Habbo API
	resp, err := p.call(ctx, url)
	if err != nil {
		return nil, err
	}

	// decode response
	var h *habbo.Habbo
	if err := json.NewDecoder(resp.Body).Decode(&h); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return h, nil
}

func (p *Parser) profile(ctx context.Context, hotel string, identifier string) (*habbo.Profile, error) {
	// create url
	var url = fmt.Sprintf("https://www.habbo.%s/api/public/users/%s/profile", hotel, identifier)

	// call Habbo API
	resp, err := p.call(ctx, url)
	if err != nil {
		return nil, err
	}

	// decode response
	var profile *habbo.Profile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return profile, nil
}

// call creates and executes a http request with the given url and user agent
func (p *Parser) call(ctx context.Context, url string) (*http.Response, error) {
	// set up request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", p.userAgent)

	// execute request
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrUnexpectedStatusCode
	}

	return resp, nil
}
