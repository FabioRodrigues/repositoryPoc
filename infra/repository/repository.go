package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vingarcia/ksql"
	"repositoryPoc/domain"
	"strings"
)

type Repository[T any] struct {
	db ksql.Provider
}

func (r Repository[T]) Save(ctx context.Context, item T) error {
	tableName := getTableName(item)

	table := ksql.NewTable(tableName)

	return r.db.Insert(ctx, table, &item)

}

func (r Repository[T]) Get(ctx context.Context, name string) (T, error) {
	var obj T

	if err := json.Unmarshal([]byte(fmt.Sprintf("{\"name\":\"%s\"}", name)), &obj); err != nil {
		return *new(T), err
	}

	return obj, nil
}

func (r Repository[T]) Transaction(ctx context.Context) {
	r.db.Transaction(ctx, func(provider ksql.Provider) error {
		return nil
	})
}

func NewRepository[T any](db ksql.Provider) domain.Repository[T] {
	return Repository[T]{
		db: db,
	}
}

func getTableName[T any](item T) string {
	typeName := fmt.Sprintf("%T", item)
	split := strings.Split(strings.ToLower(typeName), ".")
	return split[len(split)-1]
}
