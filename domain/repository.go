package domain

import (
	"context"
	"github.com/vingarcia/ksql"
)

type Repository[T any] interface {
	Save(ctx context.Context, db ksql.Provider, item T) error
	Get(ctx context.Context, db ksql.Provider, id string) (T, error)
}
