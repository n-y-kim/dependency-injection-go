package developer

type BeDeveloper struct {
	Language   string
	Experience int32
}

// List of supported backend languages
var backendLanguages = map[string]bool{
	"Golang":  true,
	"Python":  true,
	"Java":    true,
	"C#":      true,
	"Node.js": true,
	"Ruby":    true,
	"PHP":     true,
}

func (bd *BeDeveloper) InitDev(language string, experience int32) {
	bd.Language = language
	bd.Experience = experience
}

func (bd BeDeveloper) Develop() bool {
	if _, exists := backendLanguages[bd.Language]; exists && bd.Experience > 1 {
		return true
	}
	return false
}
