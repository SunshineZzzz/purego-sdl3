package main

import "github.com/SunshineZzzz/purego-sdl3/sdl"

func main() {
	defer sdl.Quit()

	var messageBox sdl.MessageBoxData
	messageBox.Flags = sdl.MessageBoxInformation
	messageBox.SetTitle("Question")
	messageBox.SetMessage("Are you sure?")

	var yesButton sdl.MessageBoxButtonData
	yesButton.Flags = sdl.MessageBoxButtonReturnKeyDefault
	yesButton.ButtonID = 1
	yesButton.SetText("Yes")

	var noButton sdl.MessageBoxButtonData
	noButton.Flags = sdl.MessageBoxButtonEscapeKeyDefault
	noButton.ButtonID = 2
	noButton.SetText("No")
	messageBox.SetButtons(yesButton, noButton)

	var pressedButton int32
	if !sdl.ShowMessageBox(&messageBox, &pressedButton) {
		sdl.LogError(sdl.LogCategoryApplication, "%s", sdl.GetError())
	}

	switch pressedButton {
	case 1:
		sdl.Log("User pressed Yes")
	case 2:
		sdl.Log("User pressed No")
	default:
		sdl.Log("User hasn't pressed any button")
	}
}
