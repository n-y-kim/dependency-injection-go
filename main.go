package main

import "dev-project/developer"

func main() {
	// This is a simple example of how to use the developer package
	// to create a side project with a backend and frontend developer.
	// The project will only be successful if both developers have
	// more than 1 year of experience and are using supported languages.
	// In this case, the project will be successful.
	beDeveloper := &developer.BeDeveloper{}
	beDeveloper.InitDev("Golang", 2)

	feDeveloper := &developer.FeDeveloper{}
	feDeveloper.InitDev("React", 2)

	sideProject := NewSideProject([]developer.Developer{beDeveloper, feDeveloper})

	if sideProject.ProjectResult() {
		println("Project successful!")
	} else {
		println("Project failed.")
	}
}
