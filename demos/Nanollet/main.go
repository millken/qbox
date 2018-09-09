package main

import (
	"log"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	rect := sciter.NewRect(100, 100, 900, 550)
	window, windowsGenerateionError := window.New(sciter.SW_MAIN|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, rect)

	if windowsGenerateionError != nil {
		log.Fatalf("Failed to generate sciter window %s", windowsGenerateionError.Error())
	}

	uiLoadingError := window.LoadFile("./02.html")
	if uiLoadingError != nil {
		log.Fatalf("Failed to load ui file %s", uiLoadingError.Error())
	}

	window.SetTitle("Hello ")
	window.Show()
	window.Run()

}
