package necessity

import (
	"os/user"
)

func MustLookupUser(username string) *user.User {
	x, err := user.Lookup(username)
	if err != nil {
		panic(err)
	}

	return x
}

func MustLookupID(uid string) *user.User {
	x, err := user.LookupId(uid)
	if err != nil {
		panic(err)
	}

	return x
}
