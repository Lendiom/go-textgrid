package textgrid

import (
	"encoding/json"
	"time"
)

//CountryCode represents one of two country short codes, US and CA
type CountryCode string

var (
	CountryCodeUnitedStates CountryCode = "US"
	CountryCodeCanada       CountryCode = "CA"
)

// TextGridTime can parse a timestamp returned in the TextGrid API and turn it into
// a valid Go Time struct.
type TextGridTime struct {
	Time  time.Time
	Valid bool
}

// NewTextGridTime returns a TextGridTime instance. val should be formatted using
// the TimeLayout.
func NewTextGridTime(val string) *TextGridTime {
	t, err := time.Parse(TimeLayout, val)
	if err == nil {
		return &TextGridTime{Time: t, Valid: true}
	} else {
		return &TextGridTime{}
	}
}

// Epoch is a time that predates the formation of the company (January 1,
// 2005). Use this for start filters when you don't want to filter old results.
var Epoch = time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)

// HeatDeath is a sentinel time that should outdate the extinction of the
// company. Use this with GetXInRange calls when you don't want to specify an
// end date. Feel free to adjust this number in the year 5960 or so.
var HeatDeath = time.Date(6000, 1, 1, 0, 0, 0, 0, time.UTC)

// The reference time, as it appears in the TextGrid API.
const TimeLayout = "Mon, 2 Jan 2006 15:04:05 -0700"

// Format expected by TextGrid for searching date ranges. Monitor and other API's
// offer better date search filters
const APISearchLayout = "2006-01-02"

func (t *TextGridTime) UnmarshalJSON(b []byte) error {
	s := new(string)
	if err := json.Unmarshal(b, s); err != nil {
		return err
	}
	if s == nil || *s == "null" || *s == "" {
		t.Valid = false
		return nil
	}
	tim, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		tim, err = time.Parse(TimeLayout, *s)
		if err != nil {
			return err
		}
	}
	*t = TextGridTime{Time: tim, Valid: true}
	return nil
}

func (t TextGridTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	b, err := json.Marshal(t.Time)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}
