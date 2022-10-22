/* File manipulation */

package core

import (
	"os"

	cp "github.com/otiai10/copy"
	"github.com/rivo/tview"
)

// Stores file to be copied
var (
	bufferPath string
	bufferName string
	bufferInfo os.FileInfo
	buffer     []byte
	cacheDir   string
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

	bufferInfo, _ = os.Lstat(bufferPath)
	if !bufferInfo.Mode().IsDir() {
		var err error
		buffer, err = os.ReadFile(bufferPath)
		if err != nil {
			panic(err)
		}
	}

	ClearCache()
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

	var err error
	if bufferInfo.Mode().IsDir() {
		err = cp.Copy(bufferPath, dest)
	} else {
		err = os.WriteFile(dest, buffer, 0755)
	}

	if err != nil {
		panic(err)
	}

	App.reloadLists()
}

// Delete's a selected file or directory, copies it to buffer
func Del(l *tview.List) {
	// copies file to buffer
	Yank(l)
	// cache directories
	if bufferInfo.Mode().IsDir() {
		var err error
		cacheDir, err = os.MkdirTemp("", "cowboyCache")
		if err != nil {
			panic(err)
		}
		err = cp.Copy(bufferPath, cacheDir)
		if err != nil {
			panic(err)
		}

	}

	os.RemoveAll(bufferPath)
	bufferPath = cacheDir
	App.reloadLists()
}

// Clear cached directories
func ClearCache() {
	if _, err := os.ReadDir(cacheDir); err == nil {
		os.RemoveAll(cacheDir)
	}
}
