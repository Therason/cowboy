package core

import (
	"github.com/rivo/tview"
	"os"
)

// Global container
var (
	App *Cowboy
)

// Current contains the current working directory, Parent contains its parent
type Cowboy struct {
	Tview *tview.Application
	Current *tview.List
	Parent *tview.List
}

// Initialize App's layout and content
func (c *Cowboy) Init() error {
	c.Tview = tview.NewApplication()

	// Get names for directories, then populate Current and Parent
	wd, pd, e := getWd()
	if e != nil {
		return e
	}
	c.Parent = NewList(pd)
	c.Current = NewList(wd)

	// Basic layout
	grid := tview.NewGrid().
		SetColumns(30, 0).
		SetRows(0).
		SetBorders(true).
		AddItem(c.Parent, 0, 0, 1, 1, 0, 0, false).
		AddItem(c.Current, 0, 1, 1, 1, 0, 0, false)

	// Controls
	c.SetHandlers()

	if err := c.Tview.SetRoot(grid, true).SetFocus(c.Current).Run(); err != nil {
		return err;
	}
	return nil;
}

// Returns current directory and parent directory
func getWd() (string, string, error) {
	currentDir, _ := os.Getwd()
	// Need to find a better way to get the parent dir
	os.Chdir("..")
	parentDir, err := os.Getwd()
	os.Chdir(currentDir)
	return currentDir, parentDir, err
}

// Generate list from given directory
func NewList(text string) *tview.List {
	list := tview.NewList()
	list.ShowSecondaryText(false).
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