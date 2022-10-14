package core

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

func (c *Cowboy) SetHandlers() {
	c.tviewHandlers()
	c.currentHandlers()
	c.parentHandlers()
}

// Overall controls
func (c *Cowboy) tviewHandlers() {
	c.Tview.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			c.Tview.Stop()
		}
		return event
	})
}

// Current list controls
func (c *Cowboy) currentHandlers() {
	c.Current.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyLeft {
			return nil
		}
		return event
	})
	c.Current.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Rune() == 'j':
			downwardNav(c.Current)
		case event.Rune() == 'k':
			c.Current.SetCurrentItem(c.Current.GetCurrentItem() - 1)
		case event.Rune() == 'h', event.Key() == tcell.KeyLeft:
			c.Tview.SetFocus(c.Parent)
			downwardNav(c.Current) // Weird way to override tview.List's default nav
		}
		return event
	})
}

// Parent list controls
func (c *Cowboy) parentHandlers() {
	c.Parent.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Rune() == 'j':
			downwardNav(c.Parent)
		case event.Rune() == 'k':
			c.Parent.SetCurrentItem(c.Parent.GetCurrentItem() - 1)
		case event.Rune() == 'l', event.Key() == tcell.KeyRight:
			c.Tview.SetFocus(c.Current)
			downwardNav(c.Parent)
		}
		return event
	})
}

// Helper function for traversing a list downwards
func downwardNav(l *tview.List) {
	if l.GetItemCount() - 1 == l.GetCurrentItem() {
		l.SetCurrentItem(0)
	} else {
		l.SetCurrentItem(l.GetCurrentItem() + 1)
	}
}