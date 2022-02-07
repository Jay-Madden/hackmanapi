package repositories

import (
	"context"
	"github.com/jackc/pgconn"
	"hackmanapi/data"
	"log"
	"time"
)

func InsertRequest(db data.Database,
	context context.Context,
	userId int,
	word string,
	length string) (int, error) {

	var comm pgconn.CommandTag
	var err error

	log.Printf("Inserting new Request with values (%v, %s, %s)\n", userId, word, length)

	if length == "" {
		comm, err = db.Pool.Exec(context,
			`INSERT INTO "Requests" ("UserId", "ReturnedWord", "Length", "Time") VALUES ($1, $2, $3, $4)`,
			userId,
			word,
			nil,
			time.Now())
	} else {
		comm, err = db.Pool.Exec(context,
			`INSERT INTO "Requests" ("UserId", "ReturnedWord", "Length", "Time") VALUES ($1, $2, $3, $4)`,
			userId,
			word,
			length,
			time.Now())
	}

	if err != nil {
		return 0, err
	}
	return int(comm.RowsAffected()), nil
}
