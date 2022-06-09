package habbo

import (
	"fmt"
	"strings"
	"time"
)

func IsValidHotel(hotel string) bool {
	hotels := []string{
		"com", "com.br", "com.tr",
		"de", "fi", "fr",
		"es", "it", "nl",
	}
	for _, h := range hotels {
		if strings.EqualFold(h, hotel) {
			return true
		}
	}
	return false
}

type Time time.Time
type RoomTime time.Time

// Time formats the Habbo API uses
// time: 2004-02-11T19:00:40.000+0000
// room: 2014-04-27T14:29:21.000+00:00
const (
	timeLayout = "2006-01-02T15:04:05.000-0700"
	roomLayout = "2006-01-02T15:04:05.000-07:00"
)

func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(timeLayout, s)
	if err != nil {
		return err
	}
	*ct = Time(nt)
	return
}

func (ct Time) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *Time) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(timeLayout))
}

func (ct *RoomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(roomLayout, s)
	if err != nil {
		return err
	}
	*ct = RoomTime(nt)
	return
}

func (ct RoomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *RoomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(roomLayout))
}

func NewTime(s string) *Time {
	t, _ := time.Parse(timeLayout, s)
	return (*Time)(&t)
}

func NewRoomTime(s string) *RoomTime {
	t, _ := time.Parse(roomLayout, s)
	return (*RoomTime)(&t)
}
