package app

import (
	"github.com/therason/cowboy/app/core"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"os"
)

func NewPrimitive(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

// Returns current directory and parent directory
func getWd() (string, string) {
	currentDir, _ := os.Getwd()
	os.Chdir("..")
	parentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Chdir(currentDir)
	return currentDir, parentDir
}

func Render() {
	core.App = &core.Cowboy {
		Tview: tview.NewApplication(),
		Page: tview.NewPages(),
	}

	//handlers
	core.App.Tview.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			core.App.Tview.Stop()
		}
		return event
	})

	// Get names for directories
	wd, pd := getWd()
	parent := NewPrimitive(pd)
	current := NewPrimitive(wd)

	grid := tview.NewGrid().
		SetColumns(30, 0).
		SetRows(0).
		SetBorders(true).
		AddItem(parent, 0, 0, 1, 1, 0, 0, false).
		AddItem(current, 0, 1, 1, 1, 0, 0, false)

	if err := core.App.Tview.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}