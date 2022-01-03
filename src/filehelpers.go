package src

import (
	cp "github.com/otiai10/copy"
	"log"
	"path"
)

func CopyRecursive(source, destination string) error {
	return cp.Copy(source, path.Join(destination, source), cp.Options{OnSymlink: func(src string) cp.SymlinkAction {
		return cp.Deep
	}})
}

func CopyNonRecursive(source, destination string) {
	log.Fatalln("non-recursive copy not yet implemented")
}
