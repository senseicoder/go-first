package main

import (
	"testing"

	github "plcoder.net/namecheck/github"
)

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := github.IsValid(tt.username)
			if got != tt.want {
				t.Errorf("IsValid() = %v, want %v for username %q", got, tt.want, tt.username)
			}
		})
	}
}
