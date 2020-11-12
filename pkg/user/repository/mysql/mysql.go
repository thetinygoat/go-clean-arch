package mysql

import (
	"database/sql"

	_user "github.com/thetinygoat/go-clean-arch/pkg/models/user"
)

// this type implements User.Repository interface
type userRepository struct {
	conn *sql.DB
}

// NewUserRepository instantiates a new user repository
func NewUserRepository(conn *sql.DB) _user.Repository {
	return &userRepository{conn: conn}
}

func (u *userRepository) GetByID(id int64) (*_user.User, error) {
	// implementation upto users
	return nil, nil
}
func (u *userRepository) Update(user *_user.User) error {
	// implementation upto users
	return nil
}
func (u *userRepository) Delete(id int64) error {
	// implementation upto users
	return nil
}
