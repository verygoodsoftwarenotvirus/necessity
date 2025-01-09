package necessity

import "mime"

func MustParseMediaType(s string) (string, map[string]string) {
	x, y, err := mime.ParseMediaType(s)
	if err != nil {
		panic(err)
	}

	return x, y
}

func MustExtensionsByType(s string) []string {
	x, err := mime.ExtensionsByType(s)
	if err != nil {
		panic(err)
	}

	return x
}
