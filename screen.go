package main

import (
	"image"
	"syscall/js"

	"github.com/fogleman/gg"
)

type Screen struct {
	Width    int
	Height   int
	Levels   []*Level
	LevelIdx int
}

func NewScreen(id string) *Screen {
	// Init Canvas stuff
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", id)
	width := int(doc.Get("body").Get("clientWidth").Float())
	height := int(doc.Get("body").Get("clientHeight").Float())
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)

	level := MakeLevel(width, height)
	return &Screen{int(width), int(height), []*Level{level}, 0}
}

func (s *Screen) Update() {
	s.Levels[s.LevelIdx].Update()
}

func (s *Screen) Draw() {

	background := image.NewRGBA(image.Rect(0, 0, s.Width, s.Height))
	c := gg.NewContextForRGBA(background)
	s.Levels[s.LevelIdx].Draw(c)

	cp := make([]byte, len(background.Pix))
	copy(cp, background.Pix)

	y := js.TypedArrayOf(cp)
	js.Global().Call("DrawClamped", s.Width, s.Height, y)
	y.Release()
}
