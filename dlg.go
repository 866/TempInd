package main

import (
	"fmt"
	"time"

	"github.com/866/tempind/read"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

// Continuously updates the label
func updateTemp(label *gtk.Label, done <-chan struct{}) {
	defer gtk.MainQuit()
EndlessLoop:
	for {
		select {
		case <-time.After(1 * time.Second):
			temp, err := read.Temp()
			if err == nil {
				text := fmt.Sprintf("%.0f °C", temp)
				gdk.ThreadsEnter()
				label.SetText(text)
				gdk.ThreadsLeave()
			} else {
				fmt.Println(err)
			}
		case <-done:
			fmt.Println("Quitting")
			break EndlessLoop
		}
	}
}

func main() {
	// Initialize threads
	glib.ThreadInit(nil)
	gdk.ThreadsInit()
	gdk.ThreadsEnter()
	gtk.Init(nil)

	// Launch a new window
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Temperature Indicator")
	window.SetSizeRequest(80, 60)
	window.SetResizable(false)
	window.SetKeepAbove(true)
	window.SetSkipTaskbarHint(true)
	fmt.Println(window.GetSkipTaskbarHint())
	done := make(chan struct{})

	// Allow the windows to be closed
	defer close(done)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "Quitting")

	// Add label that is updated by goroutine
	label := gtk.NewLabel("0 °C")
	go updateTemp(label, done)
	window.Add(label)
	window.ShowAll()

	// Run main gtk gui loop
	gtk.Main()
}
