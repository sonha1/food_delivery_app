package repo

import (
	"context"
	"github.com/sonha1/food_delivery_app/global"
	"github.com/sonha1/food_delivery_app/internal/database"
)

type UserRepo struct {
	q       *database.Queries
	context context.Context
}

func NewUserRepo(q *database.Queries) *UserRepo {
	return &UserRepo{q: q, context: context.Background()}
}

func (r *UserRepo) FindUserByUsername(username string) (*database.User, error) {
	user, err := r.q.FindUserByUsername(r.context, username)
	if err != nil {
		global.Logger.Error("FindUserByUsername: " + err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetAllUsers() (*[]database.User, error) {
	users, err := r.q.GetAllUser(r.context)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *UserRepo) CreateUser(dto *database.CreateUserParams) (*database.User, error) {
	result, err := r.q.CreateUser(r.context, *dto)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := r.q.FindUserById(r.context, int32(id))
	if err != nil {
		return nil, err
	}
	return &user, nil
}
