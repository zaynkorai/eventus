package event

import (
	"time"

	"github.com/labstack/echo"

	"github.com/zaynkorai/eventus"
	"github.com/zaynkorai/eventus/pkg/api/event"
)

// New creates new event logging service
func New(svc event.Service, logger eventus.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents event logging service
type LogService struct {
	event.Service
	logger eventus.Logger
}

const name = "Event Service"

// Create logging
func (ls *LogService) Create(c echo.Context, req eventus.Event) (resp eventus.Event, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Create event request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// List logging
func (ls *LogService) List(c echo.Context) (resp []eventus.Event, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List events request", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c)
}

func (ls *LogService) Read(c echo.Context, req int) (resp eventus.Event, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Read a event request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Read(c, req)
}

func (ls *LogService) Update(c echo.Context, req event.Update) (resp eventus.Event, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update event request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}

func (ls *LogService) Delete(c echo.Context, req int) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete event request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, req)
}
