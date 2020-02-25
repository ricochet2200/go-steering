package main

import (
	"time"
)

const MsPerFrame = 1000 / 10

func main() {

	screen := NewScreen("mycanvas")

	tick := time.Tick(MsPerFrame * time.Millisecond)

	for {
		<-tick
		screen.Update()
		screen.Draw()
	}
}
