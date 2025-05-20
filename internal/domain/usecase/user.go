package usecase

import "sample-server/internal/domain/entity"

// UserUsecaseの例

type UserUsecase interface {
	GetUser(id int) (*entity.User, error)
}

type userUsecase struct{}

func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

func (u *userUsecase) GetUser(id int) (*entity.User, error) {
	if id == 1 {
		return &entity.User{ID: 1, Name: "Ken"}, nil
	}

	return nil, entity.ErrInvalidToken
}
