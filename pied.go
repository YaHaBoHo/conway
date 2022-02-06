package main

import (
	"conway/conway"
	"conway/utilities"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

const WorldSize = 512
const CellSize float64 = 2
const WinSize = float64(WorldSize) * CellSize

func drawCell(imd *imdraw.IMDraw, x float64, y float64) {
	imd.Push(pixel.V(x, y))
	imd.Push(pixel.V(x+CellSize, y+CellSize))
	imd.Rectangle(0)
}

func drawWorld(window *pixelgl.Window, world *conway.World) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 1, 0)
	for row := 0; row < world.Size; row++ {
		for col := 0; col < world.Size; col++ {
			if world.Grid[row*world.Size+col]&1 == 1 {
				drawCell(imd, float64(row)*CellSize, float64(col)*CellSize)
			}
		}
	}
	imd.Draw(window)

}

func run() {

	frames := 0
	second := time.Tick(time.Second)

	world := conway.RandomWorld(WorldSize, 0.25)

	cfg := pixelgl.WindowConfig{
		Title:  "Conway",
		Bounds: pixel.R(0, 0, WinSize, WinSize),
		// VSync:  true,
	}

	window, err := pixelgl.NewWindow(cfg)
	utilities.CheckError(err)

	// Simulate
	for !window.Closed() {
		window.Clear(colornames.Black)
		drawWorld(window, world)
		window.Update()
		frames++
		select {
		case <-second:
			window.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
		conway.UpdateGridConcurrently(world, 6)
	}

}

func main() {
	pixelgl.Run(run)
}
