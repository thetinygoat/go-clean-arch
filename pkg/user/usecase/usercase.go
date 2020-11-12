package usecase

import _user "github.com/thetinygoat/go-clean-arch/pkg/models/user"

// this type implements User.Usecase
type userUsecase struct {
	repo *_user.Repository
}

// NewUserUsecase instantiates a new user usecase
func NewUserUsecase(ur *_user.Repository) _user.Usecase {
	return &userUsecase{repo: ur}
}

func (u *userUsecase) GetByID(id int64) (*_user.User, error) {
	// implementation upto users
	return nil, nil
}
func (u *userUsecase) Update(user *_user.User) error {
	// implementation upto users
	return nil
}
func (u *userUsecase) Delete(id int64) error {
	// implementation upto users
	return nil
}
