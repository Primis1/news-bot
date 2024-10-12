package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	model "news-bot/internal/model"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}

func (s *SourcePostgresStorage) Source(ctx context.Context) ([]model.Source, error) {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

}
func (s *SourcePostgresStorage) SourceByID(ctx context.Context, id int64) (*model.Source, error) {}
func (s *SourcePostgresStorage) Add(ctx context.Context, source model.Source) (int64, error)     {}
func (s *SourcePostgresStorage) Delete(ctx context.Context, id int64) error                      {}

type dbSource struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	FeedUrl   string `db:"feed_url"`
	AddedAt   string `db:"added_at"`
	CreatedAt string `db:"created_at"`
}
