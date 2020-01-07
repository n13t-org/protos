package gitlab

import (
	"regexp"
	"testing"
)

func TestActionTypes(t *testing.T) {
	var pattern = regexp.MustCompile("^(created|updated|closed|reopened|pushed|commented|merged|joined|left|destroyed|expired)$")
	var tests = []struct{
		text string
		want bool
	}{
		{"created", true},
		{"", false},
		{" created ", false},
		{"ccreated", false},
		{"ccreated", false},
		{"createdupdated", false},
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

func TestTargetTypes(t *testing.T) {
	var pattern = regexp.MustCompile("^(issue|milestone|merge_request|note|project|snippet|user|)$")
	var tests = []struct{
		text string
		want bool
	}{
		{"", true},
		{"issue", true},
		{" created ", false},
		{"ccreated", false},
		{"ccreated", false},
		{"createdupdated", false},
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
