package habbo

type Habbo struct {
	UniqueID                    string  `json:"uniqueId"`
	Name                        string  `json:"name"`
	FigureString                string  `json:"figureString"`
	Motto                       string  `json:"motto"`
	Online                      bool    `json:"online"`
	LastAccessTime              *Time   `json:"lastAccessTime"`
	MemberSince                 Time    `json:"memberSince"`
	ProfileVisible              bool    `json:"profileVisible"`
	CurrentLevel                int     `json:"currentLevel"`
	CurrentLevelCompletePercent int     `json:"currentLevelCompletePercent"`
	TotalExperience             int     `json:"totalExperience"`
	StarGemCount                int     `json:"starGemCount"`
	SelectedBadges              []Badge `json:"selectedBadges"`
}
