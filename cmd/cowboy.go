package cmd

import (
	"github.com/rivo/tview"
)

func Render() {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Cowboy")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}