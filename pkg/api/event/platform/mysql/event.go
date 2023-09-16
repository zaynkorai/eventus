package mysql

import (
	"errors"

	"github.com/zaynkorai/eventus"
	"gorm.io/gorm"
)

// Event represents the client for eve nt table
type Event struct{}

// Create creates a new event database
func (eve Event) Create(db *gorm.DB, event eventus.Event) (eventus.Event, error) {
	db.Create(&event)
	return event, nil
}

// Read returns single event by ID
func (eve Event) Read(db *gorm.DB, id int) (eventus.Event, error) {
	var event eventus.Event

	result := db.First(&event, "id = ?", id)

	if err := result.Error; err != nil {
		return eventus.Event{}, errors.New("No Event with that ID exists")
	}

	return event, nil
}

func (eve Event) Update(db *gorm.DB, event eventus.Event) error {
	db.Model(&event).Updates(event)
	return nil
}

// List returns list of all events
func (eve Event) List(db *gorm.DB) ([]eventus.Event, error) {
	var events []eventus.Event
	db.Find(&events)
	return events, nil
}

// Delete sets deleted_at for a event
func (eve Event) Delete(db *gorm.DB, event eventus.Event) error {
	result := db.Delete(&event, "id = ?", event.ID)

	if result.RowsAffected == 0 {
		return errors.New("No note with that Id exists")
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}
