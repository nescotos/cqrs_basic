package database

import (
	"context"
	"database/sql"

	"github.com/nestor94/cqrs_basic/model"
)

type PostgresRepository struct {
	db *sql.DB
}

func CreateConnection(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

// Close : Closes the connection with PG
func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

// Insert : Allows to insert a new Tweet
func (repo *PostgresRepository) Insert(ctx context.Context, tweet model.Tweet) error {
	_, err := repo.db.Exec(
		"INSERT INTO \"tweets\"(body) VALUES($1)",
		tweet.Body,
	)
	return err
}

// Read: Reads Tweets from DB, support for pagination using skip and take
func (repo *PostgresRepository) Read(ctx context.Context, skip uint64, take uint64) ([]model.Tweet, error) {
	rows, err := repo.db.Query(
		"SELECT * FROM \"tweets\" ORDER BY \"created_at\" ASC OFFSET $1 LIMIT $2",
		skip, take,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tweets := []model.Tweet{}
	for rows.Next() {
		tweet := model.Tweet{}
		if err = rows.Scan(&tweet.ID, &tweet.Body, &tweet.CreatedAt, &tweet.UpdatedAt, &tweet.DeletedAt); err == nil {
			tweets = append(tweets, tweet)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tweets, nil
}
