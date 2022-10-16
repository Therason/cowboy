/* Main storage struct and app initialization */

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
	Tview   *tview.Application
	Grid    *tview.Grid
	Current *tview.List
	Parent  *tview.List
}

// Initialize App's layout and content
func (c *Cowboy) Init() error {
	c.Tview = tview.NewApplication()

	// Get names for directories, then populate Current and Parent
	wd, pd, e := GetDirs()
	if e != nil {
		return e
	}
	c.Parent = NewList(pd)
	c.Current = NewList(wd)

	// Basic layout
	c.Grid = tview.NewGrid().
		SetColumns(30, 0).
		SetRows(0).
		AddItem(c.Parent, 0, 0, 1, 1, 0, 0, false).
		AddItem(c.Current, 0, 1, 1, 1, 0, 0, false)

	// Controls
	c.SetHandlers()

	if err := c.Tview.SetRoot(c.Grid, true).SetFocus(c.Current).Run(); err != nil {
		return err
	}
	return nil
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
