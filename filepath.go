package necessity

import (
	"path/filepath"
)

func MustAbs(path string) string {
	x, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return x
}

func MustRel(basepath, targpath string) string {
	x, err := filepath.Rel(basepath, targpath)
	if err != nil {
		panic(err)
	}

	return x
}
