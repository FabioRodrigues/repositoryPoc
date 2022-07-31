package services

import (
	"context"
	"github.com/vingarcia/ksql"
	"repositoryPoc/domain"
	"repositoryPoc/domain/entities"
	"repositoryPoc/infra/uow"
)

type UserService struct {
	userRepository    domain.Repository[entities.User]
	addressRepository domain.Repository[entities.Address]
	uow               uow.Uow
}

func NewUserService(
	uow uow.Uow,
	userRepository domain.Repository[entities.User],
	addressRepository domain.Repository[entities.Address],
) UserService {
	return UserService{
		addressRepository: addressRepository,
		userRepository:    userRepository,
		uow:               uow,
	}
}

func (s UserService) CreateUser(ctx context.Context, userName string, addressName string, street string) (entities.User, error) {

	var user entities.User

	txErr := s.uow.Transaction(ctx, func(provider ksql.Provider) error {
		u := entities.User{
			Id:   "123",
			Name: userName,
		}

		a := entities.Address{
			Id:     "321",
			Name:   addressName,
			Street: street,
			UserId: u.Id,
		}

		err := s.userRepository.Save(ctx, provider, u)
		if err != nil {
			return err
		}

		err = s.addressRepository.Save(ctx, provider, a)
		if err != nil {
			return err
		}
		user = u

		return nil
	})

	if txErr != nil {
		return user, txErr
	}

	return user, nil
}

func (s UserService) GetUser(ctx context.Context, id string) (entities.User, error) {
	u, err := s.userRepository.Get(ctx, s.uow.Db(), id)
	if err != nil {
		return entities.User{}, err
	}

	return u, nil

}
