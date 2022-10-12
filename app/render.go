package app

import (
	"github.com/therason/cowboy/app/core"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"os"
)

func NewList(text string) *tview.List {
	list := tview.NewList()
	list.
		ShowSecondaryText(false).
		SetTitle(text).
		SetBorder(true)


	files, err := os.ReadDir(text)
	if err != nil {
		panic(err)
	}
	for _, i := range files {
		list.AddItem(i.Name(), "", 0, nil)
	}
	return list
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
	core.App.Parent = NewList(pd)
	core.App.Current = NewList(wd)

	grid := tview.NewGrid().
		SetColumns(30, 0).
		SetRows(0).
		SetBorders(true).
		AddItem(core.App.Parent, 0, 0, 1, 1, 0, 0, false).
		AddItem(core.App.Current, 0, 1, 1, 1, 0, 0, false)

	if err := core.App.Tview.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}