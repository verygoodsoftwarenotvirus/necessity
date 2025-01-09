package necessity

import "database/sql"

func MustOpen(driverName, dataSourceName string) *sql.DB {
	x, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	return x
}
