package main

import "dev-project/developer"

func SideProject(be developer.BeDeveloper, fe developer.FeDeveloper) bool {
	return be.Develop() && fe.Develop()
}
