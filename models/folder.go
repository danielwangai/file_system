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
	err := validateFolderRootName(f.Name)
	if err != nil {
		return nil, err
	}
	f.ID = uuid.Must(uuid.NewV4()).String()
	filingSystem = append(filingSystem, f)
	return f, nil
}

// CreateSubFolder adds a new folder within an existing folder
func (f *Folder) CreateSubFolder(name string) (*Folder, error) {
	err := f.validateFolder(name)
	if err != nil {
		return nil, err
	}
	newFolder := &Folder{
		ID:     uuid.Must(uuid.NewV4()).String(),
		Name:   name,
		Parent: f,
	}
	f.Children = append(f.Children, newFolder)
	return newFolder, nil
}

// GetRootFolders gets folders at the root directory
func GetRootFolders() ([]*Folder, error) {
	if len(filingSystem) == 0 {
		return []*Folder{}, nil
	}
	root := []*Folder{}
	for _, folder := range filingSystem {
		if folder.Parent == nil {
			root = append(root, folder)
		}
	}
	return root, nil
}

// GetSubFolders get folders within a folder
func (f *Folder) GetSubFolders() ([]*Folder, error) {
	if len(f.Children) == 0 {
		return []*Folder{}, nil
	}
	return f.Children, nil
}

// UpdateFolder updates a folder's name
func (f *Folder) UpdateFolder(newName string) (*Folder, error) {
	if newName == "" {
		return nil, errors.New("provide a valid folder name")
	}
	f.Name = newName
	return f, nil
}

// MoveFolder moves folder from one location to another
func (f *Folder) MoveFolder(destFolder *Folder) {
	/*
		TODO:
			- reject moving a folder if user has insufficient permissions
			- prevent moving to a folder where user has no permission
	*/
	parent := f.Parent
	// delete folder from the existing directory
	parent.deleteFolderHelper(f)
	// move to new folder
	destFolder.Children = append(destFolder.Children, f)
}

// deleteFolderHelper deletes sub-folder within a folder
func (f *Folder) deleteFolderHelper(deletingFolder *Folder) {
	for i, subFolder := range f.Children {
		if deletingFolder == subFolder {
			f.Children = append(f.Children[:i], f.Children[i+1:]...)
			break
		}
	}
}

// validateFolderRootName ensures that a folder name is unique within the current folder
// TODO: join folder name validators into one function
func validateFolderRootName(name string) error {
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

func (f Folder) validateFolder(name string) error {
	if name == "" {
		return errors.New("folder name required")
	}
	for _, folder := range f.Children {
		if folder.Name == name {
			return errors.New("a folder with the same name exists")
		}
	}
	return nil
}
