package eventus

import "time"

type Event struct {
	ID          int       `gorm:"primary_key;auto_increment;not_null" json:"id,omitempty"`
	Title       string    `gorm:"type:varchar(255)" json:"title,omitempty"`
	Description string    `gorm:"type:varchar(255)" json:"description,omitempty"`
	StartTime   time.Time `gorm:"type:datetime(3)" json:"startTime"`
	EndTime     time.Time `gorm:"type:datetime(3)" json:"endTime"`
}
