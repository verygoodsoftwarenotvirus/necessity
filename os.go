package necessity

import "os"

func MustStat(name string) os.FileInfo {
	x, err := os.Stat(name)
	if err != nil {
		panic(err)
	}

	return x
}
