package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/demouth/ebitenlg"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	gui  *ebitenlg.GUI
	gui2 *ebitenlg.GUI
}

func (g *Game) Update() error {
	g.gui.Update()
	g.gui2.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x66, 0x66, 0x66, 0xff})
	g.gui.Draw(screen)
	g.gui2.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{}

	// gui1
	gui := ebitenlg.NewGUI()
	gui.X = screenWidth
	gui.HorizontalAlign = ebitenlg.HorizontalAlignRight
	gui.AddSliderFloat64("float64", 0, -500, 500, func(v float64) {})
	gui.AddSliderInt("int", 5, 1, 10, func(v int) {})
	gui.AddButton("bool", true, func(v bool) {})
	game.gui = gui

	// gui2
	gui2 := ebitenlg.NewGUI()
	gui2.Width = 400
	gui2.AddSliderFloat64("Scale", float64(gui.Scale), 0.1, 2, func(v float64) {
		gui.Scale = float32(v)
	})
	gui2.AddSliderFloat64("X", float64(gui.X), 0, 1000, func(v float64) {
		gui.X = float32(v)
	})
	gui2.AddSliderFloat64("Y", float64(gui.Y), 0, 1000, func(v float64) {
		gui.Y = float32(v)
	})
	gui2.AddSliderFloat64("Width", float64(gui.Width), 10, 1000, func(v float64) {
		gui.Width = float32(v)
	})
	gui2.AddSliderFloat32("ComponentHeight", gui.ComponentHeight, 5, 100, func(v float32) {
		gui.ComponentHeight = v
	})
	gui2.AddButton("Set HorizontalAlign to Right", true, func(v bool) {
		if v {
			gui.HorizontalAlign = ebitenlg.HorizontalAlignRight
		} else {
			gui.HorizontalAlign = ebitenlg.HorizontalAlignLeft
		}
	})
	game.gui2 = gui2

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine lil GUI")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
