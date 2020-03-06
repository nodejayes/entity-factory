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
		return "", errors.New("error on get current working directory: " + cwdErr.Error())
	}
	filePath := path.Join(cwd, file)
	if _, statErr := os.Stat(filePath); os.IsNotExist(statErr) {
		return "", errors.New("file does not exists: " + statErr.Error())
	}
	return filePath, nil
}

// returns the JSON Schema Directory and creates it when not exists
func (pf *Pathfinder) GetJsonSchemaFolder() (string, error) {
	cwd, cwdErr := os.Getwd()
	if cwdErr != nil {
		return "", errors.New("error on get current working directory: " + cwdErr.Error())
	}
	directoryPath := path.Join(cwd, "_schema")
	if _, statErr := os.Stat(directoryPath); os.IsNotExist(statErr) {
		createDirectoryErr := os.MkdirAll(directoryPath, 0644)
		if createDirectoryErr != nil {
			return "", errors.New("error on create directory " + directoryPath + ": " + createDirectoryErr.Error())
		}
	}
	return directoryPath, nil
}
