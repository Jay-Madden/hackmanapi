package repositories

import (
	"context"
	"hackmanapi/data"
)

func InsertRequest(db data.Database, userId int, word string) (int, error) {
	comm, err := db.Pool.Exec(context.Background(),
		`INSERT INTO "Requests" ("UserId", "Word") VALUES ($1, $2)`,
		userId,
		word)

	if err != nil {
		return 0, err
	}
	return int(comm.RowsAffected()), nil
}
