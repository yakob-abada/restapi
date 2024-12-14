package auth

import "github.com/stretchr/testify/mock"

type AuthorizationMock struct {
	mock.Mock
}

func (a *AuthorizationMock) UserId() int {
	args := a.Called()
	return args.Int(0)
}
