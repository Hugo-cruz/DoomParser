package domain

import (
	"io"
	"os"
)

var NewFs NewFileSystem = osFS{}

type NewFileSystem interface {
	Open(name string) (file, error)
	Stat(name string) (os.FileInfo, error)
}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

type osFS struct{}

func (osFS) Open(name string) (file, error)        { return os.Open(name) }
func (osFS) Stat(name string) (os.FileInfo, error) { return os.Stat(name) }

type MockedFS struct {
	osFS
	reportErr bool
}

type MockedFile struct {
	os.FileInfo
}

func (m MockedFile) Open(name string) (file, error) { return os.Open(name) }
