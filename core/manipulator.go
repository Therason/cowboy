/* File manipulation */

package core

import (
	"github.com/rivo/tview"
	cp "github.com/otiai10/copy"
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


	// Copy src buffer to destination
	currentPath, parentPath, _ := GetDirs()
	var dest string
	if l == App.Current {
		dest = currentPath + "/" + bufferName
	} else {
		dest = parentPath + "/" + bufferName
	}

	srcInfo, err := os.Lstat(bufferPath)
	if err != nil {
		panic(err)
	}

	if srcInfo.Mode().IsDir() {
		err = cp.Copy(bufferPath, dest)
	} else {
		// Set source file buffer
		src, err := os.ReadFile(bufferPath)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(dest, src, 0755)
	}


	if err != nil {
		panic(err)
	}

	App.reloadLists()
}
