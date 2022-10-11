package core

import (
	"github.com/rivo/tview"
)

var (
	App *Cowboy
)

type Cowboy struct {
	Tview *tview.Application
	Page *tview.Pages
}