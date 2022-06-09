package habbo

type Badge struct {
	BadgeIndex  int    `json:"badgeIndex"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
