/* Directory navigation */

package core

import (
	"os"
)

// Returns current directory and parent directory
func GetDirs() (string, string, error) {
	currentDir, _ := os.Getwd()
	// Need to find a better way to get the parent dir
	os.Chdir("..")
	parentDir, err := os.Getwd()
	os.Chdir(currentDir)
	return currentDir, parentDir, err
}

// Navigate into selected directory
func (c *Cowboy) TraverseDirDown() {
	targetDir, _ := c.Current.GetItemText(c.Current.GetCurrentItem())
	if err := os.Chdir(targetDir); err != nil {
		// Not a directory, try to display the file
		loadFile(targetDir)
		return
	}
	c.Current.Clear()
	c.Parent.Clear()

	c.reloadLists()
}

// Navigate to parent directory
func (c *Cowboy) TraverseDirUp() {
	// TODO: when parent traverses up, set current item to parent's parent
	c.Current.Clear()
	c.Parent.Clear()
	if e := os.Chdir(".."); e != nil {
		panic(e)
	}

	c.reloadLists()
}

// Allows parent directory to change child directory
func (c *Cowboy) ParentTraverseDirDown() {
	targetDir, _ := c.Parent.GetItemText(c.Parent.GetCurrentItem())
	currentDir, _ := c.Current.GetItemText(c.Current.GetCurrentItem())
	if targetDir != currentDir {
		if err := os.Chdir("../" + targetDir); err != nil {
			// Not a directory
			return
		}
		c.reloadLists()
	}
	c.Tview.SetFocus(c.Current)
}

// Helper function to fill Current and Parent with proper items
func (c *Cowboy) reloadLists() {
	c.Current.Clear()
	c.Parent.Clear()
	current, parent, err := GetDirs()
	if err != nil {
		panic(err)
	}

	c.Current.SetTitle(current)
	c.Parent.SetTitle(parent)

	currentFiles, _ := os.ReadDir(current)
	parentFiles, _ := os.ReadDir(parent)

	for _, i := range currentFiles {
		c.Current.AddItem(i.Name(), "", 0, nil)
	}
	for _, i := range parentFiles {
		c.Parent.AddItem(i.Name(), "", 0, nil)
	}
}

// Read data from file
func loadFile(target string) {
	// Read data from target
	output, err := os.ReadFile(target)
	if err != nil {
		panic(err)
	}

	// TODO: format data into current grid panel
	//fmt.Println(string(output[:]))
	App.OpenFile.SetText(string(output[:])).SetBorder(true)
	App.Grid.RemoveItem(App.Current).
		AddItem(App.OpenFile, 0, 1, 1, 1, 0, 0, false)
}
