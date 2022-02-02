package repositories

import (
	"context"
	"hackmanapi/data"
)

func InsertUser(db data.Database, name string, key string) (int, error) {
	comm, err := db.Pool.Exec(context.Background(),
		`INSERT INTO "Requests" ("Name", "ApiKey") VALUES ($1, $2)`,
		name,
		key)

	if err != nil {
		return 0, err
	}
	return int(comm.RowsAffected()), nil
}
