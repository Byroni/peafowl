package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/byroni/peafowl/ports"
)

func New(service ports.Peafowl) {
	a := app.New()
	w := a.NewWindow("Hello world")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("welcome")
		}),
	))

	w.ShowAndRun()
}
