package main

import (
	"dev-project/developer"
	"testing"
)

func TestSideProject(t *testing.T) {
	tests := []struct {
		beDev    developer.BeDeveloper
		feDev    developer.FeDeveloper
		expected bool
	}{
		{developer.BeDeveloper{Language: "Golang", Experience: 2}, developer.FeDeveloper{Language: "JavaScript", Experience: 2}, true},
		{developer.BeDeveloper{Language: "Python", Experience: 1}, developer.FeDeveloper{Language: "JavaScript", Experience: 2}, false},
	}

	for _, tt := range tests {
		got := SideProject(tt.beDev, tt.feDev)
		if got != tt.expected {
			t.Errorf("SideProject(%v, %v) = %v; want %v", tt.beDev, tt.feDev, got, tt.expected)
		}
	}
}
