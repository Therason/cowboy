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

func (c *Cowboy) TraverseDirDown() {
	targetDir, _ := c.Current.GetItemText(c.Current.GetCurrentItem())
	if err := os.Chdir(targetDir); err != nil {
		// Not a directory
		return
	}
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

func (c *Cowboy) TraverseDirUp() error {
	c.Current.Clear()
	c.Parent.Clear()
	if e := os.Chdir(".."); e != nil {
		return e
	}

	current, parent, err := GetDirs()
	if err != nil {
		return err
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

	return nil
}
