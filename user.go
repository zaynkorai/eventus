package eventus

import "time"

// User represents user domain model
type User struct {
	ID        int    `gorm:"primary_key;auto_increment;not_null" json:"id,omitempty"`
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	LoginType string `json:"login_type"`

	Code     int    `json:"-"`
	Token    string `json:"-"`
	Password string `json:"-"`

	LastLogin time.Time `json:"last_login,omitempty"`
}

// UpdateLastLogin updates last login field
func (u *User) UpdateLastLogin(token string) {
	u.Token = token
	u.LastLogin = time.Now()
}
