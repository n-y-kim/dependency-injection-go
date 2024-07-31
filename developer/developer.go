package developer

type Developer interface {
	InitDev(language string, experience int32)
	Develop() bool
}
