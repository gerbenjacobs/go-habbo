package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gerbenjacobs/go-habbo/habbo"
)

type ParserMock struct {
	HabboResponse   *habbo.Habbo
	HabboError      error
	ProfileResponse *habbo.Profile
	ProfileError    error
}

func NewParserMock() *ParserMock {
	return &ParserMock{}
}

func (p *ParserMock) habbo(context.Context, string, string, bool) (*habbo.Habbo, error) {
	return p.HabboResponse, p.HabboError
}

func (p *ParserMock) profile(context.Context, string, string) (*habbo.Profile, error) {
	return p.ProfileResponse, p.ProfileError
}

func (p *ParserMock) loadHabbo(file string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read Habbo file: %v", err)
	}
	if err := json.Unmarshal(bytes, &p.HabboResponse); err != nil {
		log.Fatalf("failed to unmarshal Habbo response: %v", err)
	}
}

func (p *ParserMock) loadProfile(file string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read Habbo file: %v", err)
	}
	if err := json.Unmarshal(bytes, &p.ProfileResponse); err != nil {
		log.Fatalf("failed to unmarshal Habbo response: %v", err)
	}
}
