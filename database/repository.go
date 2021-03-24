package database

import (
	"context"

	"github.com/nestor94/cqrs_basic/model"
)

// TweetRepository : Interface to define the behaviour of an implementation for TweetRepository regardless Database
type TweetRepository interface {
	Close()
	Insert(ctx context.Context, tweet model.Tweet) error
	Read(ctx context.Context, skip uint64, take uint64) ([]model.Tweet, error)
}

var implementation TweetRepository

// SetRepository : Allows to set the implementation for TweetRepository
func SetRepository(tweetRepository TweetRepository) {
	implementation = tweetRepository
}

// Close : Close a connection to the database
func Close() {
	implementation.Close()
}

// Insert : Allows to insert a new Tweet into the Database
func Insert(ctx context.Context, tweet model.Tweet) error {
	return implementation.Insert(ctx, tweet)
}

// Read: Allows to read from the Database, supports pagination using skip and take parameters
func Read(ctx context.Context, skip uint64, take uint64) ([]model.Tweet, error) {
	return implementation.Read(ctx, skip, take)
}
