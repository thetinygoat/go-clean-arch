// Package http contains http handlers for the user
// it uses usecase methods to do the operations
package http

import _user "github.com/thetinygoat/go-clean-arch/pkg/models/user"

type userDelivery struct {
	usecase *_user.Usecase
}

// NewUserHTTPDelivery registers users routes and returns the route handler to be used when starting ther server
func NewUserHTTPDelivery(us *_user.Usecase) {
	// register user routes
}
