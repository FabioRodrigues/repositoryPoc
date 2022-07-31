package repository

import (
	"context"
	"fmt"
	"github.com/vingarcia/ksql"
	"repositoryPoc/domain"
	"strings"
)

type Repository[T any] struct {
}

func (r Repository[T]) Save(ctx context.Context, db ksql.Provider, item T) error {
	tableName := getTableName(item)

	table := ksql.NewTable(tableName)

	return db.Insert(ctx, table, &item)

}

func (r Repository[T]) Get(ctx context.Context, db ksql.Provider, id string) (T, error) {

	var obj T

	tableName := getTableName(obj)
	query := fmt.Sprintf("FROM \"%s\" WHERE id = $1", tableName)

	err := db.QueryOne(ctx, &obj, query, id)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

func NewRepository[T any]() domain.Repository[T] {
	return Repository[T]{}
}

func getTableName[T any](item T) string {
	typeName := fmt.Sprintf("%T", item)
	split := strings.Split(strings.ToLower(typeName), ".")
	return split[len(split)-1]
}
