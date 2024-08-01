package main

import "dev-project/developer"

type SideProject struct {
	developers []developer.Developer
}

func (s SideProject) ProjectResult() bool {
	for _, developer := range s.developers {
		if !developer.Develop() {
			return false
		}
	}
	return true
}
