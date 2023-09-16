package eventus

import (
	"net/http"

	"errors"

	"github.com/labstack/echo"
)

var (
	// ErrGeneric is used for testing purposes and for errors handled later in the callstack
	ErrGeneric = errors.New("Generic error")

	// ErrBadRequest (400) is returned for bad request (validation)
	ErrBadRequest = echo.NewHTTPError(400)

	// ErrUnauthorized (401) is returned when user is not authorized
	ErrUnauthorized = echo.ErrUnauthorized

	// ErrBadRequest (400) is returned for when data is not abailable
	ErrNoData = echo.NewHTTPError(http.StatusBadRequest, "No data available")

	ErrInvalidPayload = echo.NewHTTPError(http.StatusBadRequest, "Error invalid data type or missing required data")
)
