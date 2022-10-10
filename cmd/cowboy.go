package cmd

import (
	"github.com/rivo/tview"
)

func NewPrimitive(text string) tview.Primitive {
	return tview.NewTextView().
					SetTextAlign(tview.AlignCenter).
					SetText(text)
}

func Render() {
	parent := NewPrimitive("Parent")
	current := NewPrimitive("Current")

	grid := tview.NewGrid().
					SetColumns(20, 0).
					SetRows(0).
					SetBorders(true).
					AddItem(parent, 0, 0, 1, 1, 0, 0, false).
					AddItem(current, 0, 1, 1, 1, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}

}