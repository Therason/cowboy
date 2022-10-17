/* Controls */

package core

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Store highlighted file state when traversing lists
var (
	currentIndex int = 0
	parentIndex  int = 0
)

// Setup controls
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
			ClearCache()
		}
		return event
	})
}

// Current list controls
func (c *Cowboy) currentHandlers() {
	c.Current.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Rune() == 'j', event.Key() == tcell.KeyDown:
			downwardNav(c.Current)
			return nil
		case event.Rune() == 'k', event.Key() == tcell.KeyUp:
			c.Current.SetCurrentItem(c.Current.GetCurrentItem() - 1)
			return nil
		case event.Rune() == 'h', event.Key() == tcell.KeyLeft:
			currentIndex = c.Current.GetCurrentItem()
			c.Tview.SetFocus(c.Parent)
			c.Current.SetCurrentItem(currentIndex)
			return nil
		case event.Rune() == 'l', event.Key() == tcell.KeyRight:
			parentIndex = c.Current.GetCurrentItem()
			c.TraverseDirDown()
			c.Parent.SetCurrentItem(parentIndex)
			return nil
		case event.Rune() == 'y':
			Yank(c.Current)
		case event.Rune() == 'p':
			Put(c.Current)
		case event.Rune() == 'd':
			Del(c.Current)
		}
		return event
	})
}

// Parent list controls
func (c *Cowboy) parentHandlers() {
	c.Parent.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Rune() == 'j', event.Key() == tcell.KeyDown:
			downwardNav(c.Parent)
			return nil
		case event.Rune() == 'k', event.Key() == tcell.KeyUp:
			c.Parent.SetCurrentItem(c.Parent.GetCurrentItem() - 1)
			return nil
		case event.Rune() == 'l', event.Key() == tcell.KeyRight:
			currentIndex = c.Current.GetCurrentItem()
			parentIndex = c.Parent.GetCurrentItem()
			c.ParentTraverseDirDown()
			c.Current.SetCurrentItem(currentIndex)
			c.Parent.SetCurrentItem(parentIndex)
			return nil
		case event.Rune() == 'h', event.Key() == tcell.KeyLeft:
			currentIndex = c.Parent.GetCurrentItem()
			c.TraverseDirUp()
			c.Current.SetCurrentItem(currentIndex)
			return nil
		case event.Rune() == 'y':
			Yank(c.Parent)
		case event.Rune() == 'p':
			Put(c.Parent)
		case event.Rune() == 'd':
			Del(c.Parent)
		}
		return event
	})
}

// Helper function for traversing down through a list
func downwardNav(l *tview.List) {
	if l.GetItemCount()-1 == l.GetCurrentItem() {
		l.SetCurrentItem(0)
	} else {
		l.SetCurrentItem(l.GetCurrentItem() + 1)
	}
}
