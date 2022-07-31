package uow

import (
	"context"
	"github.com/vingarcia/ksql"
)

type Uow interface {
	Transaction(ctx context.Context, f func(provider ksql.Provider) error) error
	Db() ksql.Provider
}

type UnitOfWork struct {
	db ksql.Provider
}

func NewUnitOfWork(db ksql.Provider) Uow {
	return UnitOfWork{
		db: db,
	}
}

func (uow UnitOfWork) Transaction(ctx context.Context, f func(provider ksql.Provider) error) error {
	return uow.db.Transaction(ctx, f)
}

// Db returns the instance of the Database in case you don't want to
// deal with transactions, when making a simple query, for instance
func (uow UnitOfWork) Db() ksql.Provider {
	return uow.db
}
