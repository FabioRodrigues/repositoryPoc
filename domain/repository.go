package domain

import "context"

type Repository[T any] interface {
	Save(ctx context.Context, item T) error
	Get(ctx context.Context, id string) (T, error)
}
