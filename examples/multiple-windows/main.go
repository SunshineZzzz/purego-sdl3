package main

import (
	"fmt"

	"github.com/SunshineZzzz/purego-sdl3/sdl"
)

var inputState = InputState{}

func main() {
	if !sdl.SetHint(sdl.HintRenderVSync, "1") {
		panic(sdl.GetError())
	}

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	// Create 3 windows and renderers
	windows := make([]Window, 0, 3)
	for i := 0; i < cap(windows); i++ {
		w := Window{}
		if !sdl.CreateWindowAndRenderer(fmt.Sprintf("Window %d", i+1), 640, 360, sdl.WindowResizable, &w.W, &w.R) {
			panic(sdl.GetError())
		}
		w.WindowID = sdl.GetWindowID(w.W)                // Get ID of the created window, to process events correctly
		sdl.SetWindowPosition(w.W, 10+650*int32(i), 100) // Position the windows next to each other
		windows = append(windows, w)
	}
	// Defer cleaning up the windows
	defer func() {
		for i := 0; i < len(windows); i++ {
			sdl.DestroyRenderer(windows[i].R)
			sdl.DestroyWindow(windows[i].W)
		}
	}()

	running := true
	for running {
		inputState.MouseLeftClicked = false // Reset left clicked state

		// Process events
		event := sdl.Event{}
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				fallthrough
			case sdl.EventWindowCloseRequested:
				running = false

			case sdl.EventMouseMotion:
				e := event.Motion()
				inputState.MouseOnWindowID = e.WindowID
				inputState.MousePosition.X, inputState.MousePosition.Y = e.X, e.Y

			case sdl.EventMouseButtonDown:
				e := event.Button()
				inputState.MouseOnWindowID = e.WindowID
				inputState.MouseLeftPressed = e.Button == uint8(sdl.ButtonLeft)
			case sdl.EventMouseButtonUp:
				e := event.Button()
				inputState.MouseOnWindowID = e.WindowID
				inputState.MouseLeftPressed = false
				inputState.MouseLeftClicked = e.Button == uint8(sdl.ButtonLeft)

			case sdl.EventWindowMouseEnter:
				e := event.Window()
				// "Autofocus" the window when mouse enters it.
				// If the window is not focused and the user clicked the button:
				// - the 1st click will focus the window and won't register button click,
				// - next clicks over the button will register button clicks.
				for i := range windows {
					w := &windows[i]
					if e.WindowID == w.WindowID {
						sdl.RaiseWindow(w.W)
						break
					}
				}
			}
		}

		// Draw contents for each of the windows
		for i := range windows {
			w := &windows[i]

			sdl.SetRenderDrawColor(w.R, 30, 30, 30, 255)
			sdl.RenderClear(w.R)

			// Draw simple button
			if DoButton(w, sdl.FRect{X: 50, Y: 100, W: 80, H: 40}) {
				w.ClickCounter++
			}
			// Draw number of times the button was clicked
			sdl.SetRenderDrawColor(w.R, 255, 255, 255, 255)
			sdl.RenderDebugText(w.R, 135, 118, fmt.Sprintf("Clicked: %d times", w.ClickCounter))

			sdl.RenderPresent(w.R)
		}
	}
}

// Structure to keep window and renderer together
type Window struct {
	W *sdl.Window   // SDL window
	R *sdl.Renderer // SDL renderer

	WindowID     sdl.WindowID
	ClickCounter int // Number of times button inside the window was clicked
}

// Input devices state
type InputState struct {
	MouseOnWindowID                    sdl.WindowID
	MousePosition                      sdl.FPoint
	MouseLeftPressed, MouseLeftClicked bool
}

// Draws a button on provided window. Returns true if the button was clicked, otherwise false
func DoButton(w *Window, rect sdl.FRect) bool {
	clicked := false
	if inputState.MouseOnWindowID == w.WindowID && sdl.PointInRectFloat(inputState.MousePosition, rect) {
		// If point (mouse position) is inside a rect (button rectangle) change drawing color
		if inputState.MouseLeftPressed {
			// If mouse left button is pressed set different color
			sdl.SetRenderDrawColor(w.R, 100, 100, 200, 255)
		} else {
			// Mouse is over the simple button but mouse left button is not pressed, draw with "button hovered" color
			sdl.SetRenderDrawColor(w.R, 100, 100, 100, 255)
			if inputState.MouseLeftClicked {
				// Mouse click detected, the simple button was pressed
				clicked = true
			}
		}
	} else {
		// Mouse is outside the button, draw the button with default color
		sdl.SetRenderDrawColor(w.R, 200, 200, 200, 255)
	}
	sdl.RenderFillRect(w.R, &rect) // Draw the button

	return clicked
}
