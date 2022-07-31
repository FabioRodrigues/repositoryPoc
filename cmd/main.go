// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
	"log"
	"repositoryPoc/domain/entities"
	"repositoryPoc/domain/services"
	"repositoryPoc/infra/repository"
	uow2 "repositoryPoc/infra/uow"
)

func main() {

	ctx := context.Background()
	dbUrl := "postgres://postgres:postgres@localhost:5432/test_repo"
	db, err := kpgx.New(ctx, dbUrl, ksql.Config{})
	if err != nil {
		log.Fatalf("unable connect to database: %s", err)
	}
	defer db.Close()

	uow := uow2.NewUnitOfWork(db)

	userRepo := repository.NewRepository[entities.User]()
	addressRepo := repository.NewRepository[entities.Address]()

	svc := services.NewUserService(uow, userRepo, addressRepo)

	createdUser, err := svc.CreateUser(ctx, "Fabio", "home", "01, homeland street")

	retrievedUser, err := svc.GetUser(ctx, "123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(createdUser)
	fmt.Println(retrievedUser)

}
