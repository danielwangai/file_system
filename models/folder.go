package models

import (
	"errors"

	"github.com/satori/go.uuid"
)

// Folder - definition of folder properties
type Folder struct {
	ID string
	// FolderType string // root or other
	Name     string
	Parent   *Folder
	Children []*Folder
}

// store folders and files
var filingSystem []*Folder

// CreateRootFolder adds a new folder to the root of the file system
func (f *Folder) CreateRootFolder() (*Folder, error) {
	err := validateFolderName(f.Name)
	if err != nil {
		return nil, err
	}
	f.ID = uuid.Must(uuid.NewV4()).String()
	filingSystem = append(filingSystem, f)
	return f, nil
}

// CreateSubFolder adds a new folder within an existing folder
func (f *Folder) CreateSubFolder(name string) (*Folder, error) {
	err := validateFolderName(name)
	if err != nil {
		return nil, err
	}
	newFolder := &Folder{
		ID:     uuid.Must(uuid.NewV4()).String(),
		Name:   name,
		Parent: f,
	}
	// the folder is being created within another one(as a sub-folder)
	f.Children = append(f.Children, newFolder)
	// parentFolder.Children = append(parentFolder.Children, f)
	return newFolder, nil
}

func validateFolderName(name string) error {
	if name == "" {
		return errors.New("folder name required")
	}
	for _, folder := range filingSystem {
		if name == folder.Name {
			return errors.New("a folder with the same name exists")
		}
	}
	return nil
}
