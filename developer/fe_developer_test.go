package developer

import "testing"

func TestFeDeveloper_Develop(t *testing.T) {
	tests := []struct {
		language   string
		experience int32
		expected   bool
	}{
		{"JavaScript", 2, true},
		{"TypeScript", 1, false},
		{"React", 3, true},
		{"Unknown", 2, false},
	}

	for _, tt := range tests {
		fd := FeDeveloper{}
		fd.InitDev(tt.language, tt.experience)
		result := fd.Develop()
		if result != tt.expected {
			t.Errorf("FeDeveloper.Develop() with language %s and experience %d = %v; want %v", tt.language, tt.experience, result, tt.expected)
		}
	}
}
