package main

import (
	"dev-project/developer"
	"testing"
)

func TestSideProject(t *testing.T) {
	beDev := &developer.BeDeveloper{Language: "Golang", Experience: 2}
	feDev := &developer.FeDeveloper{Language: "JavaScript", Experience: 2}
	project := NewSideProject([]developer.Developer{beDev, feDev})

	if !project.ProjectResult() {
		t.Errorf("Expected project to succeed, but it failed")
	}

	beDev = &developer.BeDeveloper{Language: "Python", Experience: 1}
	project = NewSideProject([]developer.Developer{beDev, feDev})

	if project.ProjectResult() {
		t.Errorf("Expected project to fail, but it succeeded")
	}
}
