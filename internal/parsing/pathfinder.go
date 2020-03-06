package parsing

import (
	"errors"
	"os"
	"path"
)

type Pathfinder struct{}

// returns the Path to a File in the Current Working Directory only when exists
func (pf *Pathfinder) GetFileInCurrentWorkingDirectory(file string) (string, error) {
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		return "", errors.New("error on get current working directory")
	}
	filePath := path.Join(cwd, file)
	if _, statErr := os.Stat(filePath); os.IsNotExist(statErr) {
		return "", errors.New("file does not exists")
	}
	return filePath, nil
}
