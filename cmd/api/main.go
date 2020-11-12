package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_userHttpDelivery "github.com/thetinygoat/go-clean-arch/pkg/user/delivery/http"
	_userRepo "github.com/thetinygoat/go-clean-arch/pkg/user/repository/mysql"
	_userUsecase "github.com/thetinygoat/go-clean-arch/pkg/user/usecase"
)

func main() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(&userRepo)
	_userHttpDelivery.NewUserHTTPDelivery(&userUsecase)
	// start server
}
