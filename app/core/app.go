package core

import (
	"github.com/rivo/tview"
)

var (
	App *Cowboy
)

type Cowboy struct {
	Tview *tview.Application
	Current *tview.List
	Parent *tview.List
}