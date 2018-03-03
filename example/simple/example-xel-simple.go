package main

import (
	"github.com/amortaza/go-xel"
	"runtime"
)

func onLoop() {
	// this gets called every loop...use it wisely
}

func onAfterLoadGL() {
	// This callback gets called after the OpenGL context has been
	// initialized.
	// All OpenGL calls are valid now
}

func onBeforeUnloadGL() {
	// Right before the window is deleted, this callback can be used
	// to free up resources.
}

func onResize(width, height int) {
	// This callback gets called for every resize of the window.
}

func onMouseMove(x, y int) {
	// This callback gets called for every movement of the mouse.
}

func onMouseButton(button xel.MouseButton, action xel.ButtonAction) {
	// This callback gets called for every press of the mouse button.
}

func onKey(key xel.KeyboardKey, action xel.ButtonAction, alt, ctrl, shift bool) {
	// This callback gets called for every press of the keyboard.
}

func init() {
	// This bit is required for GLFW.
	runtime.LockOSThread()
}

func main() {

	// Create an 800 by 600 pixel window (is this 1997 again!)
	xel.Init(800, 600)

	// Setup the callbacks
	xel.SetCallbacks(onAfterLoadGL, onLoop, onBeforeUnloadGL, onResize, onMouseMove, onMouseButton, onKey)

	// Start the loop!
	xel.Loop("Hello, World!")
}