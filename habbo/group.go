package habbo

type Group struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Type            string `json:"type"`
	RoomID          string `json:"roomId"`
	BadgeCode       string `json:"badgeCode"`
	PrimaryColour   string `json:"primaryColour"`
	SecondaryColour string `json:"secondaryColour"`
	Online          bool   `json:"online"`
	IsAdmin         bool   `json:"isAdmin"`
}
