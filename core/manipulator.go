package core

import (
	"github.com/rivo/tview"
	"os"
)

// Stores file to be copied
var (
	bufferPath string
	bufferName string
)

// Sets buffer to selected file path
func Yank(l *tview.List) {
	bufferName, _ = l.GetItemText(l.GetCurrentItem())
	currentPath, parentPath, _ := GetDirs()
	// TODO: list container struct that stores the list and the containter's working dir
	if l == App.Current {
		bufferPath = currentPath + "/" + bufferName
	} else {
		bufferPath = parentPath + "/" + bufferName
	}
}

// Paste a file
func Put(l *tview.List) {
	// Set source file buffer
	src, err := os.ReadFile(bufferPath)
	if err != nil {
		panic(err)
	}

	// Copy src buffer to destination
	currentPath, parentPath, _ := GetDirs()
	if l == App.Current {
		dest := currentPath + "/" + bufferName
		err = os.WriteFile(dest, src, 0755)
	} else {
		dest := parentPath + "/" + bufferName
		err = os.WriteFile(dest, src, 0755)
	}

	if err != nil {
		panic(err)
	}

	App.reloadLists()
}
