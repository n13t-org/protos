package test

import (
	"regexp"
	"testing"
)

func TestZoneInfoPattern(t *testing.T) {
	var pattern = regexp.MustCompile("^\\w+(/\\w+){1,2}([-|+][0-9]{1,2})?$")
	var tests = []struct{
		text string
		want bool
	}{
		{"asdf", false},
		{"asdf2", false},
		{"America/North_Dakota/New_Salem", true},
		{"America/Matamoros", true},
		{"America/Matamoros/", false},
		{"America//", false},
		{"America/", false},
		{"America", false},
		{"Etc/GMT", true},
		{"Etc/GMT-10", true},
		{"Etc/GMT+10", true},
		{"Etc/GMT-14", true},
		{"Etc/GMT+12", true},
		//{"Etc/GMT-15", false}, // skip! not support yet
		//{"Etc/GMT+13", false}, // skip! not support yet
		{"Etc/UTC", true},
	}

	for _, tt := range tests {
		t.Run("\""+tt.text+"\"", func(t *testing.T) {
			ans := pattern.MatchString(tt.text)
			if !(ans == tt.want) {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}

func TestEmailPattern(t *testing.T) {
	var pattern = regexp.MustCompile("^\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")
	var tests = []struct{
		text string
		want bool
	}{
		{"asdf", false},
		{"a@a.c", true},
		{"user@example.com", true},
		{"user@example", false},
		{"@example", false},
		{"user@", false},
		{"user@@", false},
		{"user@exmaple..com", false},
	}

	for _, tt := range tests {
		t.Run("\""+tt.text+"\"", func(t *testing.T) {
			ans := pattern.MatchString(tt.text)
			if !(ans == tt.want) {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
