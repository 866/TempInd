package main

import (
	"fmt"
	"time"

	"tempind/read"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdk"
)

// Continuously updates the label
func updateTemp(label *gtk.Label, done <-chan struct{}) {
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
	gtk.Init(nil)
	dialog := gtk.NewWindow(0)
	dialog.SetTitle("Temperature Indicator")
	dialog.SetSizeRequest(80, 60)
	dialog.SetResizable(false)
	done := make(chan struct{})
	dialog.Connect("destroy", func(ctx *glib.CallbackContext) {
		close(done)
		gtk.MainQuit()
	}, "Quitting")
	label := gtk.NewLabel("0 °C")
	go updateTemp(label, done)
	dialog.Add(label)
	dialog.ShowAll()
	gtk.Main()
}
