package core

import (
	"github.com/rivo/tview"
	"os"
)

var (
	App *Cowboy
)

type Cowboy struct {
	Tview *tview.Application
	Current *tview.List
	Parent *tview.List
}

func (c *Cowboy) Init() error {
	c.Tview = tview.NewApplication()


	// Get names for directories
	wd, pd := getWd()
	c.Parent = NewList(pd)
	c.Current = NewList(wd)

	grid := tview.NewGrid().
		SetColumns(30, 0).
		SetRows(0).
		SetBorders(true).
		AddItem(c.Parent, 0, 0, 1, 1, 0, 0, false).
		AddItem(c.Current, 0, 1, 1, 1, 0, 0, false)

	c.SetHandlers()

	if err := c.Tview.SetRoot(grid, true).SetFocus(c.Current).Run(); err != nil {
		return err;
	}
	return nil;
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