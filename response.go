package worldtime

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	EXPECTED_VERSION = "1.2"
)

type Response struct {
	Current struct {
		Abbreviation   string `json:"abbreviation"`
		Description    string `json:"description"`
		EffectiveUntil string `json:"effectiveUntil"`
		IsDst          bool   `json:"isDst"`
		UtcOffset      string `json:"utcOffset"`
	} `json:"current"`

	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Region    string  `json:"region"`
	} `json:"location"`

	Next struct {
		Abbreviation   string `json:"abbreviation"`
		Description    string `json:"description"`
		EffectiveUntil string `json:"effectiveUntil"`
		IsDst          bool   `json:"isDst"`
		UtcOffset      string `json:"utcOffset"`
	} `json:"next"`

	Summary struct {
		HasDst bool   `json:"hasDst"`
		Local  string `json:"local"`
		Utc    string `json:"utc"`
	} `json:"summary"`

	URL     string `json:"url"`
	Version string `json:"version"`
}

/////////////////////////////////////////////////
// Methods
/////////////////////////////////////////////////

func (r *Response) CurrentUTCOffset() float64 {

	input := r.Current.UtcOffset
	result := 0.0

	positive := true

	if string(input[0]) == "-" {
		positive = false
	}

	remainder := string(input[1:])

	parts := strings.Split(remainder, ":")

	if hour, h_err := strconv.ParseInt(parts[0], 10, 0); h_err == nil {
		result = float64(hour)
	}

	if minute, m_err := strconv.ParseInt(parts[1], 10, 0); m_err == nil {
		result += float64(minute) / 60
	}

	if !positive {
		result *= -1
	}

	return result
}

/////////////////////////////////////////////////
// Utility functions
/////////////////////////////////////////////////

func NewResponse() *Response {
	return &Response{}
}

func ParseResponse(input []byte) (r *Response, e error) {

	r = NewResponse()

	e = json.Unmarshal(input, r)

	if e != nil {
		return
	}

	if r.Version != EXPECTED_VERSION {
		return nil, fmt.Errorf("Expected version %s, got %s", EXPECTED_VERSION, r.Version)
	}

	return
}
