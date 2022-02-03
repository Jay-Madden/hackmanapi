package repositories

import (
	"context"
	"github.com/jackc/pgconn"
	"hackmanapi/data"
)

func InsertRequest(db data.Database, userId int, word string, length string) (int, error) {

	var comm pgconn.CommandTag
	var err error
	if length == "" {
		comm, err = db.Pool.Exec(context.Background(),
			`INSERT INTO "Requests" ("UserId", "ReturnedWord", "Length") VALUES ($1, $2, $3)`,
			userId,
			word,
			nil)
	} else {
		comm, err = db.Pool.Exec(context.Background(),
			`INSERT INTO "Requests" ("UserId", "ReturnedWord", "Length") VALUES ($1, $2, $3)`,
			userId,
			word,
			length)
	}

	if err != nil {
		return 0, err
	}
	return int(comm.RowsAffected()), nil
}
