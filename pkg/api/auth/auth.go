package auth

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/zaynkorai/eventus"
)

// Custom errors
var (
	ErrInvalidCredentials = echo.NewHTTPError(http.StatusUnauthorized, "Username or password does not exist")
)

// Create creates a new user account
func (a Auth) CreateUser(c echo.Context, req eventus.User) (eventus.User, error) {

	req.Password = a.sec.Hash(req.Password)
	usr, err := a.udb.Create(a.db, req)
	if err != nil {
		return eventus.User{}, err
	}

	return usr, err
}

// Authenticate tries to authenticate the user provided by username and password
func (a Auth) Authenticate(c echo.Context, user, pass string) (eventus.AuthToken, error) {
	u, err := a.udb.FindByUsername(a.db, user)
	if err != nil {
		return eventus.AuthToken{}, err
	}

	if !a.sec.HashMatchesPassword(u.Password, pass) {
		return eventus.AuthToken{}, ErrInvalidCredentials
	}

	token, err := a.tg.GenerateToken(u)
	if err != nil {
		return eventus.AuthToken{}, eventus.ErrUnauthorized
	}

	u.UpdateLastLogin(a.sec.Token(token))

	u.LoginType = "email"
	if err := a.udb.Update(a.db, u); err != nil {
		return eventus.AuthToken{}, err
	}

	return eventus.AuthToken{Token: token, RefreshToken: u.Token}, nil
}

// Refresh refreshes jwt token and puts new claims inside
func (a Auth) Refresh(c echo.Context, refreshToken string) (string, error) {
	user, err := a.udb.FindByToken(a.db, refreshToken)
	if err != nil {
		return "", err
	}
	return a.tg.GenerateToken(user)
}
