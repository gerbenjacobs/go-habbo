package habbo

type Profile struct {
	Habbo   Habbo   `json:"user"`
	Badges  []Badge `json:"badges"`
	Groups  []Group `json:"groups"`
	Friends []Habbo `json:"friends"`
	Rooms   []Room  `json:"rooms"`
}
