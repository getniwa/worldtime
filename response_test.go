package worldtime

import "testing"

func Test_CreateResponse(t *testing.T) {
	NewResponse()
}

func Test_ParseResponse(t *testing.T) {

	jsn := `{
  "version": "1.2",
  "url": "http://worldtime.io/current/WzUwLjgyNDYsIC0wLjE1NTQ5OTk5OTk5OTk5XXwtMC4xNTU0OTk5OTk5OTk5OXw1MC44MjQ2",
  "location": {
    "region": "United Kingdom",
    "latitude": 50.8246,
    "longitude": -0.15549999999999
  },
  "summary": {
    "utc": "2015-06-14 13:28:14",
    "local": "2015-06-14 14:28:14",
    "hasDst": true
  },
  "current": {
    "abbreviation": "BST",
    "description": "British Summer Time",
    "utcOffset": "+1:00",
    "isDst": true,
    "effectiveUntil": "2015-10-25 02:00:00"
  },
  "next": {
    "abbreviation": "GMT",
    "description": "Greenwich Mean Time",
    "utcOffset": "+0:00",
    "isDst": false,
    "effectiveUntil": "2016-03-27 01:00:00"
  }
}`

	_, e := ParseResponse([]byte(jsn))

	if e != nil {
		t.Errorf("Unexpected error: %s", e)
	}
}

func Test_UTCOffsets(t *testing.T) {

	offsets := []struct {
		Input  string
		Output float64
	}{
		{"+1:00", 1},
		{"+1:30", 1.5},
		{"-7:30", -7.5},
		{"-12:30", -12.5},
	}

	for _, offset := range offsets {

		r := NewResponse()

		r.Current.UtcOffset = offset.Input

		if g, e := r.CurrentUTCOffset(), offset.Output; g != e {
			t.Fatalf("[%s] got %f, expected %f", offset.Input, g, e)
		}
	}
}
