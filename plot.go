package main

import (
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func errorSplashScreen(message string) {
	drv := fyne.CurrentApp().Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		w := drv.CreateSplashWindow()
		w.SetContent(widget.NewLabelWithStyle(message,
			fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
		w.Show()

		go func() {
			time.Sleep(time.Second * 2)
			w.Close()
		}()
	}
}

func onPlotButtonPressed() {
	errorSplashScreen("success!")
}

func onEnterPressed() {
	errorSplashScreen("enter pressed!")
}

func main() {
	a := app.NewWithID("plot")
	w := a.NewWindow("plot")
	w.Resize(fyne.NewSize(200, 300))
	fEntry := NewCustomEntry(onEnterPressed)
	fEntry.SetPlaceHolder("e.g. f(x)=2x+3")
	fText := canvas.NewText("Enter Function: ", color.White)
	menu := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, fText, nil), fText, fEntry)
	plotContent := NewPlotWindow()
	plotButton := widget.NewButton("plot!", onPlotButtonPressed)
	w.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(menu, plotButton, nil, nil), menu, plotContent, plotButton))

	w.ShowAndRun()
}
