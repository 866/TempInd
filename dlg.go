package main

import (
	"fmt"
	"time"

	"github.com/866/TempInd/read"
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
	dialog := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	dialog.SetTitle("Temperature Indicator")
	dialog.SetSizeRequest(80, 60)
	dialog.SetResizable(false)
	dialog.SetKeepAbove(true)
	dialog.SetSkipTaskbarHint(true)	
	done := make(chan struct{})

	// Allow the windows to be closed
	defer close(done)
	dialog.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "Quitting")

	// Add label that is updated by goroutine
	label := gtk.NewLabel("0 °C")
	go updateTemp(label, done)
	dialog.Add(label)
	dialog.ShowAll()

	// Run main gtk gui loop
	gtk.Main()
}
