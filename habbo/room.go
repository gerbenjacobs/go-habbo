package habbo

type Room struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	CreationTime    *RoomTime `json:"creationTime"`
	Tags            []string  `json:"tags"`
	MaximumVisitors int       `json:"maximumVisitors"`
	ShowOwnerName   bool      `json:"showOwnerName"`
	OwnerName       string    `json:"ownerName"`
	OwnerUniqueId   string    `json:"ownerUniqueId"`
	Categories      []string  `json:"categories"`
	ThumbnailURL    string    `json:"thumbnailUrl"`
	ImageURL        string    `json:"imageUrl"`
	Rating          int       `json:"rating"`
	UniqueID        string    `json:"uniqueId"`
	HabboGroupID    string    `json:"habboGroupId"`
	PublicRoom      bool      `json:"publicRoom"`
	DoorMode        string    `json:"doorMode"`
}
