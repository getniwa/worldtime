package worldtime

import "testing"

const (
	TEST_KEY = "TOKEN"
)

func Test_SetKey(t *testing.T) {
	SetMashapeKey(TEST_KEY)
}

func Test_Request(t *testing.T) {
	_, err := Request("86.27.87.99")

	if err != nil {
		t.Errorf("Request: %s", err)
	}
}
