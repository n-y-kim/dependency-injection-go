package main

import (
	"dev-project/developer"
	"fmt"
)

func main() {
	beDev := &developer.BeDeveloper{}
	feDev := &developer.FeDeveloper{}

	beDev.InitDev("Python", 2)
	feDev.InitDev("React", 3)

	if SideProject(*beDev, *feDev) {
		fmt.Println("Successfully finished the project!")
	} else {
		fmt.Println("Failed the project. :(")
	}
}
