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
)

func main() {

	ctx := context.Background()
	dbUrl := "postgres://postgres:postgres@localhost:5432/test_repo"
	db, err := kpgx.New(ctx, dbUrl, ksql.Config{})
	if err != nil {
		log.Fatalf("unable connect to database: %s", err)
	}
	defer db.Close()

	userRepo := repository.NewRepository[entities.User](db)
	addressRepo := repository.NewRepository[entities.Address](db)

	svc := services.NewUserService(userRepo, addressRepo)
	u := svc.CreateUser(ctx, "Fabio", "home", "01, homeland street")

	fmt.Println(u)

}
