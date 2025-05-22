package bluesky

import (
	"testing"
)

func TestIsValidViaSlice(t *testing.T) {
	type TestCase struct {
		desc     string
		username string
		want     bool
	}
	testCases := []TestCase{
		{"contains two consecutive hyphens", "jub0bs--on-GitHub", false},
		// other test cases...
	}

	var b Bluesky

	for _, tc := range testCases {
		got, _ := b.IsValid(tc.username)
		if got != tc.want {
			const tmpl = "%s: github.IsValid(%q): got %t; want %t"
			t.Errorf(tmpl, tc.desc, tc.username, got, tc.want)
		}
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		username string
		want     bool
	}{
		{
			name:     "starts with hyphen",
			username: "-cedric",
			want:     false,
		},
		{
			name:     "double hyphen",
			username: "cedric--test",
			want:     false,
		},
		{
			name:     "too short",
			username: "ce",
			want:     false,
		},
		{
			name:     "too long",
			username: "cedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedriccedric",
			want:     false,
		},
		{
			name:     "valid username",
			username: "cedric",
			want:     true,
		},
	}

	var b Bluesky

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := b.IsValid(tt.username)
			if got != tt.want {
				t.Errorf("IsValid() = %v, want %v for username %q", got, tt.want, tt.username)
			}
		})
	}
}
