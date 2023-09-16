package transport

import (
	"net/http"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zaynkorai/eventus"
	"github.com/zaynkorai/eventus/pkg/api/event"

	"github.com/labstack/echo"
)

type HTTP struct {
	svc event.Service
}

// NewHTTP creates new event http service
func NewHTTP(svc event.Service, r *echo.Group) {
	h := HTTP{svc}
	ur := r.Group("/events")
	ur.POST("", h.create)
	ur.GET("", h.list)
	ur.GET("/:id", h.view)
	ur.PATCH("/:id", h.update)
	ur.DELETE("/:id", h.delete)
}

type createReq struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

func (h HTTP) create(c echo.Context) error {
	req := new(createReq)

	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("Bind Error: ")
		return eventus.ErrInvalidPayload
	}

	providr, err := h.svc.Create(c, eventus.Event{
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, providr)
}

type listResponse struct {
	Event []eventus.Event `json:"events"`
}

func (h HTTP) list(c echo.Context) error {

	result, err := h.svc.List(c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result})
}

func (h HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return eventus.ErrBadRequest
	}

	result, err := h.svc.Read(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

type updateReq struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (h HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return eventus.ErrBadRequest
	}

	req := new(updateReq)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("Bind Error: ")
		return eventus.ErrInvalidPayload
	}

	providr, err := h.svc.Update(c, event.Update{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, providr)
}

func (h HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return eventus.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
