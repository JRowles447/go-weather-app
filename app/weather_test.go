package app

import (
	"fmt"
	"testing"
)

// const invalidZipsDigits := []string{"12", "557852", "0", "69621475"}
// const invalidZipsChars := []string{"foo", "another", ""}

func TestValidateZip(t *testing.T) {
	var tests = []struct {
		zip  string
		want bool
	}{
		{"80209", true},
		{"98109", true},
		{"02210", true},
		{"02144", true},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.zip)

		t.Run(testName, func(t *testing.T) {
			res := ValidateZip(tt.zip)
			if res != tt.want {
				t.Errorf("got %t, want %t", res, tt.want)
			}
		})
	}
}
