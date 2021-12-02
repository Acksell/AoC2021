package util

import (
	"bufio"
	"os"
)

// A Loadable defines a load function that takes a string and outputs a potential error.
// Used for loading inputs into memory.
type Loadable interface {
	Load(string) error
}

// ReadLines reads and loads a file line by line located at `path`.
func ReadLines(path string, l Loadable) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		err := l.Load(txt)
		if err != nil {
			return err
		}
	}
	return nil
}
