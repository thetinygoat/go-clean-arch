package user

// User struct holds user blueprint
type User struct {
	Name  string
	Email string
	ID    int64
}

// Repository is the user repository interface
type Repository interface {
	GetByID(id int64) (*User, error)
	Update(u *User) error
	Delete(id int64) error
}

// Usecase is the user usecase interface
type Usecase interface {
	GetByID(id int64) (*User, error)
	Update(u *User) error
	Delete(id int64) error
}
