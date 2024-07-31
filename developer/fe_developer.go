package developer

type FeDeveloper struct {
	Language   string
	Experience int32
}

// List of supported frontend languages
var frontendLanguages = map[string]bool{
	"JavaScript": true,
	"TypeScript": true,
	"React":      true,
	"Vue.js":     true,
	"Angular":    true,
}

func (fd *FeDeveloper) InitDev(language string, experience int32) {
	fd.Language = language
	fd.Experience = experience
}

func (fd FeDeveloper) Develop() bool {
	if _, exists := frontendLanguages[fd.Language]; exists && fd.Experience > 1 {
		return true
	}
	return false
}
