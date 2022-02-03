package repositories

import (
	"context"
	"hackmanapi/data"
	"hackmanapi/data/models"
)

func InsertUser(db data.Database, name string, key string) (int, error) {
	comm, err := db.Pool.Exec(context.Background(),
		`INSERT INTO "Users" ("Name", "ApiKey") VALUES ($1, $2)`,
		name,
		key)

	if err != nil {
		return 0, err
	}
	return int(comm.RowsAffected()), nil
}

func GetUserByKey(db data.Database, key string) (models.User, error) {
	user := models.User{}
	err := db.Pool.QueryRow(context.Background(),
		`SELECT * FROM "Users" WHERE "ApiKey" = $1`, key).Scan(&user.Id, &user.Name, &user.ApiKey)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
