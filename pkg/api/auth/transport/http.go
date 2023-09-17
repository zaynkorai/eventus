package transport

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zaynkorai/eventus"
	"github.com/zaynkorai/eventus/pkg/api/auth"

	"github.com/labstack/echo"
)

// HTTP represents auth http service
type HTTP struct {
	svc auth.Service
}

// NewHTTP creates new auth http service
func NewHTTP(svc auth.Service, e *echo.Echo, mw echo.MiddlewareFunc) {
	h := HTTP{svc}

	e.POST("/signup", h.signup)
	e.POST("/login", h.login)

	e.GET("/refresh/:token", h.refresh)
}

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "Passwords do not match")
	ErrAuthFailed          = echo.NewHTTPError(http.StatusBadRequest, "Username or password does not matched")
	ErrSessionExpired      = echo.NewHTTPError(http.StatusBadRequest, "Session expired please login again")
)

// Account create request
type signupReq struct {
	FullName string `json:"full_name"`
	Username string `json:"username" validate:"required,min=3,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

func (h HTTP) signup(c echo.Context) error {
	user := new(signupReq)

	if err := c.Bind(user); err != nil {
		log.Error().Err(err).Msg("Bind Error: ")
		return eventus.ErrInvalidPayload
	}

	usr, err := h.svc.CreateUser(c, eventus.User{
		FullName: user.FullName,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

type credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *HTTP) login(c echo.Context) error {
	cred := new(credentials)
	if err := c.Bind(cred); err != nil {
		log.Error().Err(err).Msg("Bind Error: ")
		return eventus.ErrInvalidPayload
	}
	r, err := h.svc.Authenticate(c, cred.Username, cred.Password)
	if err != nil {
		return ErrAuthFailed
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) refresh(c echo.Context) error {
	token, err := h.svc.Refresh(c, c.Param("token"))
	if err != nil {
		return ErrSessionExpired
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
