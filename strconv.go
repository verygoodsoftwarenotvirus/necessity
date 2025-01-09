package necessity

import (
	"strconv"
)

func MustAtoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return x
}

func MustParseInt(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return x
}

func MustParseUint(s string) uint64 {
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return x
}

func MustParseFloat(s string) float64 {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return x
}

func MustParseBool(s string) bool {
	x, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}

	return x
}
