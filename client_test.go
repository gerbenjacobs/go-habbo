package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/gerbenjacobs/go-habbo/habbo"
)

var (
	habboKoeientemmer = &habbo.Habbo{
		UniqueID:       "hhus-9cd61b156972c2eb33a145d69918f965",
		Name:           "koeientemmer",
		FigureString:   "hd-195-1.ch-215-75.lg-3290-91.sh-905-1408.ha-1015.wa-2001",
		Motto:          "Oldskooler than Dionysus!",
		MemberSince:    *habbo.NewTime("2001-10-06T12:21:53.000+0000"),
		ProfileVisible: true,
		SelectedBadges: []habbo.Badge{
			{
				BadgeIndex:  1,
				Code:        "YODUK",
				Name:        "Master Yoduck",
				Description: "Use the nondescript motion imparting field!",
			},
			{
				BadgeIndex:  2,
				Code:        "UK183",
				Name:        "The Sims - Katy Perry Sweet Treats",
				Description: "I beat the The Sims Katy Perry Sweet Treats Quiz!",
			},
			{
				BadgeIndex:  3,
				Code:        "Z63_HHUK",
				Name:        "Valued BETA tester",
				Description: "Helped shape the new Habbo June 2009",
			},

			{
				BadgeIndex:  4,
				Code:        "Z64_HHUK",
				Name:        "Official BETA tester",
				Description: "Helped shape the new Habbo June 2009",
			},
			{
				BadgeIndex:  5,
				Code:        "UK119",
				Name:        "Master Shifu's Badge of Honour",
				Description: "Kung Fu Panda 2 visited Habbo December 2010",
			},
		},
	}
	habboJohno = &habbo.Habbo{
		UniqueID:                    "hhs2-15cdd228b60baf1fcd72283ab29d1527",
		Name:                        "Johno",
		FigureString:                "hr-165-31.hd-180-1.ch-215-1408.lg-280-1408.sh-300-1408.fa-1201",
		Motto:                       "null",
		Online:                      false,
		LastAccessTime:              habbo.NewTime("2020-12-07T15:43:37.000+0000"),
		MemberSince:                 *habbo.NewTime("2004-07-19T10:14:20.000+0000"),
		ProfileVisible:              true,
		CurrentLevel:                7,
		CurrentLevelCompletePercent: 75,
		TotalExperience:             110,
		StarGemCount:                18,
		SelectedBadges: []habbo.Badge{
			{
				BadgeIndex:  2,
				Code:        "ADM",
				Name:        "",
				Description: "",
			},
			{
				BadgeIndex:  4,
				Code:        "ACH_HappyHour1",
				Name:        "Happy hour",
				Description: "For logging in during happy hour.",
			},
			{
				BadgeIndex:  5,
				Code:        "ACH_AllTimeHotelPresence7",
				Name:        "Online time VII - Cyclone",
				Description: "For spending total of 1440  hours in hotel.",
			},
		},
	}
	habboGerben = &habbo.Habbo{
		UniqueID:                    "hhnl-d47ce47a95c35b2d3f027d207e3d0515",
		Name:                        "Gerben",
		FigureString:                "hr-802-36.hd-195-1.ch-3030-64.lg-3088-81-80.sh-290-64.ha-3454.cc-3389-64-90",
		Motto:                       "Duikt de kast weer in!",
		Online:                      false,
		LastAccessTime:              habbo.NewTime("2020-12-31T13:50:01.000+0000"),
		MemberSince:                 *habbo.NewTime("2004-02-11T19:00:40.000+0000"),
		ProfileVisible:              true,
		CurrentLevel:                26,
		CurrentLevelCompletePercent: 18,
		TotalExperience:             1429,
		StarGemCount:                60,
		SelectedBadges: []habbo.Badge{
			{
				BadgeIndex:  1,
				Code:        "APC13",
				Name:        "Ze houden je in de gaten...maar hebben je niet gezien!",
				Description: "Habbocalyps 2015",
			},
			{
				BadgeIndex:  2,
				Code:        "SB7",
				Name:        "Streets of Bobba - 1",
				Description: "Eerste plaats",
			},
		},
	}
)

func TestGetHabbo_ValidateUniqueID(t *testing.T) {
	tests := []struct {
		name      string
		hotel     string
		habboName string
		wantErr   bool
		errorType error
	}{
		{name: "valid", hotel: "com", habboName: "hhus-9cd61b156972c2eb33a145d69918f965", wantErr: false},
		{name: "invalid hotel", hotel: "invalid", habboName: "hhus-9cd61b156972c2eb33a145d69918f965", wantErr: true, errorType: ErrInvalidHotel},
		{name: "invalid unique ID", hotel: "com", habboName: "invalid", wantErr: true, errorType: ErrInvalidUniqueID},
		{name: "sandbox ID", hotel: "com", habboName: "hhs2-15cdd228b60baf1fcd72283ab29d1527", wantErr: false},
	}
	for _, tt := range tests {
		api := &HabboAPI{
			parser: NewParserMock(),
		}
		t.Run(tt.name, func(t *testing.T) {
			_, err := api.GetHabbo(context.Background(), tt.hotel, tt.habboName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHabboByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errorType.Error() {
				t.Errorf("GetHabboByName() error = %v, wantErr %v", err, tt.errorType)
			}
		})
	}
}

func TestGetHabbo_Parsing(t *testing.T) {
	tests := []struct {
		name    string
		hotel   string
		habboID string
		file    string
		want    *habbo.Habbo
		wantErr bool
	}{
		{
			name:    "valid for koeientemmer",
			hotel:   "com",
			habboID: "hhus-9cd61b156972c2eb33a145d69918f965",
			file:    "tests/data/com_koeientemmer_gethabbo.json",
			want:    habboKoeientemmer,
			wantErr: false,
		},
		{
			name:    "valid sandbox hotel for johno",
			hotel:   "com", // not 100%, but it's good enough for this test
			habboID: "hhs2-15cdd228b60baf1fcd72283ab29d1527",
			file:    "tests/data/sandbox_johno_gethabbo.json",
			want:    habboJohno,
			wantErr: false,
		},
		{
			name:    "valid for gerben",
			hotel:   "nl",
			habboID: "hhnl-d47ce47a95c35b2d3f027d207e3d0515",
			file:    "tests/data/nl_gerben_gethabbo.json",
			want:    habboGerben,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		pm := NewParserMock()
		api := &HabboAPI{
			parser: pm,
		}
		t.Run(tt.name, func(t *testing.T) {
			pm.loadHabbo(tt.file)
			got, err := api.GetHabbo(context.Background(), tt.hotel, tt.habboID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHabbo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHabbo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProfile_Koeientemmer(t *testing.T) {
	pm := NewParserMock()
	pm.loadProfile("tests/data/com_koeientemmer_getprofile.json")
	api := &HabboAPI{
		parser: pm,
	}

	p, err := api.GetProfile(context.Background(), "com", "hhus-9cd61b156972c2eb33a145d69918f965")
	if err != nil {
		t.Error(err)
	}

	// check user entity
	if !reflect.DeepEqual(p.Habbo, *habboKoeientemmer) {
		t.Errorf("GetProfile() got = \n%#v, want \n%#v", p.Habbo, *habboKoeientemmer)
	}

	// check badges
	if len(p.Badges) != 204 {
		t.Errorf("GetProfile() got = %v, want %v", len(p.Badges), 204)
	}
	expectedBadge := habbo.Badge{
		Code:        "THI41",
		Name:        "...will go on!!!",
		Description: "Winner of broken hearts 2/2",
	}
	if !reflect.DeepEqual(p.Badges[0], expectedBadge) {
		t.Errorf("GetProfile() got = \n%#v, want \n%#v", p.Badges[0], expectedBadge)
	}

	// check rooms
	if len(p.Rooms) != 5 {
		t.Errorf("GetProfile() got = %v, want %v", len(p.Rooms), 5)
	}
	expectedRoom := habbo.Room{
		ID:              31159787,
		Name:            "Venice Beach Rollercoaster",
		Description:     "Ahh well..",
		CreationTime:    habbo.NewRoomTime("2010-06-10T09:02:16.000+00:00"),
		MaximumVisitors: 25,
		Tags: []string{
			"habbies",
		},
		ShowOwnerName: true,
		OwnerName:     "koeientemmer",
		OwnerUniqueId: "hhus-9cd61b156972c2eb33a145d69918f965",
		ThumbnailURL:  "https://habbo-stories-content.s3.amazonaws.com/navigator-thumbnail/hhus/31159787.png",
		ImageURL:      "https://habbo-stories-content.s3.amazonaws.com/fullroom-photo/hhus/31159787.png",
		Rating:        116,
		Categories: []string{
			"navigator.flatcategory.global.CHAT",
		},
		UniqueID: "r-hhus-a091e0f1d891108b49ca7af953386f0f",
	}
	if !reflect.DeepEqual(p.Rooms[0], expectedRoom) {
		t.Errorf("GetProfile() got = \n%#v, want \n%#v", p.Rooms[0], expectedRoom)
	}

	// check groups
	if len(p.Groups) != 10 {
		t.Errorf("GetProfile() got = %v, want %v", len(p.Groups), 10)
	}
	expectedGroup := habbo.Group{
		ID:              "g-hhus-1332dcd15645042afc396f726351721d",
		Name:            "Ditch the Label",
		Description:     "The official Ditch the Label anti-bullying public group. Join now and support our cause! Find out more at www.DitchtheLabel.org",
		Type:            "NORMAL",
		RoomID:          "r-hhus-a08de337a9dc601102b0139194164f78",
		BadgeCode:       "b13114s19134a55aa7427bc0a3f0c083e94232fb3475",
		PrimaryColour:   "242424",
		SecondaryColour: "ffffff",
		IsAdmin:         false,
	}
	if !reflect.DeepEqual(p.Groups[0], expectedGroup) {
		t.Errorf("GetProfile() got = \n%#v, want \n%#v", p.Groups[0], expectedGroup)
	}

	// check friends
	if len(p.Friends) != 146 {
		t.Errorf("GetProfile() got = %v, want %v", len(p.Friends), 146)
	}
	expectedFriend := habbo.Habbo{
		Name:         "!!!toon!!!",
		Motto:        "Puhekupla.com",
		UniqueID:     "hhus-eafa7f30310f817172231ac9ba3a3baa",
		FigureString: "hr-155-48.hd-180-1.ch-215-1408.lg-3023-64.sh-3068-1408-85.wa-2001.cp-3286",
	}
	if !reflect.DeepEqual(p.Friends[0], expectedFriend) {
		t.Errorf("GetProfile() got = \n%#v, want \n%#v", p.Friends[0], expectedFriend)
	}
}

func TestGetProfile_Johno(t *testing.T) {
	pm := NewParserMock()
	pm.loadProfile("tests/data/sandbox_johno_getprofile.json")
	api := &HabboAPI{
		parser: pm,
	}

	p, err := api.GetProfile(context.Background(), "com", "hhs2-15cdd228b60baf1fcd72283ab29d1527")
	if err != nil {
		t.Error(err)
	}

	// check user entity
	if !reflect.DeepEqual(p.Habbo, *habboJohno) {
		t.Errorf("GetProfile() got = \n%#v, want \n%#v", p.Habbo, *habboJohno)
	}

	// check counts
	if len(p.Friends) != 14 {
		t.Errorf("Friends got = %v, want %v", len(p.Friends), 14)
	}
	if len(p.Rooms) != 2 {
		t.Errorf("Rooms got = %v, want %v", len(p.Rooms), 2)
	}
	if len(p.Groups) != 0 {
		t.Errorf("Groups got = %v, want %v", len(p.Groups), 0)
	}
	if len(p.Badges) != 33 {
		t.Errorf("Badges got = %v, want %v", len(p.Badges), 33)
	}
}
