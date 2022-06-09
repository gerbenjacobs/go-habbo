package client

import (
	"context"
	"regexp"

	"github.com/gerbenjacobs/go-habbo/habbo"
)

var uniqueIdRegexp = regexp.MustCompile(`^hh\w{2}-[a-zA-Z\d]{32}$`)

// HabboParser is an interface that the Habbo API uses to fetch data.
type HabboParser interface {
	habbo(ctx context.Context, hotel string, habboID string, byName bool) (*habbo.Habbo, error)
	profile(ctx context.Context, hotel string, habboID string) (*habbo.Profile, error)
}

// HabboAPI is a Habbo API client.
type HabboAPI struct {
	parser HabboParser
}

// NewHabboAPI creates a new Habbo API client.
func NewHabboAPI(parser HabboParser) *HabboAPI {
	return &HabboAPI{parser: parser}
}

// GetHabbo fetches a Habbo by ID.
func (c *HabboAPI) GetHabbo(ctx context.Context, hotel string, habboID string) (*habbo.Habbo, error) {
	if !habbo.IsValidHotel(hotel) {
		return nil, ErrInvalidHotel
	}
	if !uniqueIdRegexp.MatchString(habboID) {
		return nil, ErrInvalidUniqueID
	}
	return c.parser.habbo(ctx, hotel, habboID, true)
}

// GetHabboByName fetches a Habbo by name.
func (c *HabboAPI) GetHabboByName(ctx context.Context, hotel string, habboName string) (*habbo.Habbo, error) {
	if !habbo.IsValidHotel(hotel) {
		return nil, ErrInvalidHotel
	}
	return c.parser.habbo(ctx, hotel, habboName, false)
}

// GetProfile fetches a Habbo's profile.
func (c *HabboAPI) GetProfile(ctx context.Context, hotel string, habboID string) (*habbo.Profile, error) {
	if !habbo.IsValidHotel(hotel) {
		return nil, ErrInvalidHotel
	}
	return c.parser.profile(ctx, hotel, habboID)
}
