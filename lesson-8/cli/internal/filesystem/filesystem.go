package filesystem

import (
	"cli/cli/internal/filesystem/mock"
	"cli/cli/internal/filesystem/real"
	model "cli/cli/internal/model/filesystem"
)

// объявляем интерфейс, с 2 методами - получить список и удалить файлы
type Filesystem interface {
	ListFiles(dirPath string) (*model.FileStats, error)
	DeleteFile(dirPath, name string) error
}

// и фукнция создать новый экземпляр файловой системы, в зависимости от типа системы
func NewFileSystem(kind string) (Filesystem, error) {
	switch kind {
	case mock.FileSystemKind:
		return mock.NewFileSystem(), nil
	case real.FileSystemKind:
		return real.NewFileSystem(), nil
	}
	return nil, model.ErrInvalidFilesystem{FilesystemKind: kind}
}
