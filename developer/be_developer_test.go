package developer

import "testing"

func TestBeDeveloper_Develop(t *testing.T) {
	tests := []struct {
		language   string
		experience int32
		expected   bool
	}{
		{"Golang", 2, true},
		{"Python", 1, false},
		{"Java", 3, true},
		{"Unknown", 2, false},
	}

	for _, tt := range tests {
		bd := BeDeveloper{}
		bd.InitDev(tt.language, tt.experience)
		result := bd.Develop()
		if result != tt.expected {
			t.Errorf("BeDeveloper.Develop() with language %s and experience %d = %v; want %v", tt.language, tt.experience, result, tt.expected)
		}
	}
}
