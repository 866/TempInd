package main

import (
	"fmt"
	"time"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"

	"tempind/read"
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
				label.SetText(text)
			} else {
				fmt.Println(err)
			}
		case <-done:
			break EndlessLoop
		}
	}
}

func main() {
	gtk.Init(nil)
	dialog := gtk.NewDialog()
	dialog.SetTitle("Temperature Indicator")
	dialog.SetDefaultSize(60, 60)
	done := make(chan struct{})
	dialog.Connect("destroy", func(ctx *glib.CallbackContext) {
		close(done)
		gtk.MainQuit()
	}, "Quitting")
	vbox := dialog.GetVBox()
	label := gtk.NewLabel("0 °C")
	go updateTemp(label, done)
	vbox.Add(label)
	dialog.ShowAll()
	gtk.Main()
}
