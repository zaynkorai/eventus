package event

import (
	"github.com/labstack/echo"

	"github.com/zaynkorai/eventus"
)

// Create creates a new event
func (eve Event) Create(c echo.Context, req eventus.Event) (eventus.Event, error) {
	return eve.edb.Create(eve.db, req)
}

// List returns list of events
func (eve Event) List(c echo.Context) ([]eventus.Event, error) {
	return eve.edb.List(eve.db)
}

// Read returns single event record
func (eve Event) Read(c echo.Context, id int) (eventus.Event, error) {
	return eve.edb.Read(eve.db, id)
}

// Update contains events's information used for updating
type Update struct {
	ID          int
	Title       string
	Description string
}

// Update updates event's contact information
func (eve Event) Update(c echo.Context, r Update) (eventus.Event, error) {
	if err := eve.edb.Update(eve.db, eventus.Event{
		ID:          r.ID,
		Title:       r.Title,
		Description: r.Description,
	}); err != nil {
		return eventus.Event{}, err
	}

	return eve.edb.Read(eve.db, r.ID)
}

// Delete deletes a event
func (eve Event) Delete(c echo.Context, id int) error {
	event, err := eve.edb.Read(eve.db, id)
	if err != nil {
		return err
	}

	return eve.edb.Delete(eve.db, event)
}
