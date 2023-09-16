package event

import (
	"github.com/zaynkorai/eventus"
	"github.com/zaynkorai/eventus/pkg/api/event/platform/mysql"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Service represents event interface
type Service interface {
	Create(echo.Context, eventus.Event) (eventus.Event, error)
	List(echo.Context) ([]eventus.Event, error)
	Read(echo.Context, int) (eventus.Event, error)
	Update(echo.Context, Update) (eventus.Event, error)
	Delete(echo.Context, int) error
}

// New creates new event service
func New(db *gorm.DB, edb EDB) *Event {
	return &Event{db: db, edb: edb}
}

// Initialize initalizes event service with defaults
func Initialize(db *gorm.DB) *Event {
	return New(db, mysql.Event{})
}

// Event represents event service
type Event struct {
	db  *gorm.DB
	edb EDB
}

// EDB represents event repository interface
type EDB interface {
	Create(*gorm.DB, eventus.Event) (eventus.Event, error)
	Read(*gorm.DB, int) (eventus.Event, error)
	List(*gorm.DB) ([]eventus.Event, error)
	Update(*gorm.DB, eventus.Event) error
	Delete(*gorm.DB, eventus.Event) error
}
