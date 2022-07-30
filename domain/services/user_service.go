package services

import (
	"context"
	"repositoryPoc/domain"
	"repositoryPoc/domain/entities"
)

type UserService struct {
	userRepository    domain.Repository[entities.User]
	addressRepository domain.Repository[entities.Address]
}

func NewUserService(
	userRepository domain.Repository[entities.User],
	addressRepository domain.Repository[entities.Address]) UserService {
	return UserService{
		addressRepository: addressRepository,
		userRepository:    userRepository,
	}
}

func (s UserService) CreateUser(ctx context.Context, userName string, addressName string, street string) entities.User {
	u := entities.User{
		Id:   "123",
		Name: userName,
	}

	err := s.userRepository.Save(ctx, u)
	if err != nil {
		panic(err)
	}

	a := entities.Address{
		Id:     "321",
		Name:   addressName,
		Street: street,
		UserId: u.Id,
	}

	err = s.addressRepository.Save(ctx, a)
	if err != nil {
		panic(err)
	}

	return u
}
