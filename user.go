package eventus

// User represents user domain model
type User struct {
	// Base
	FullName  string `json:"full_name"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	LoginType string `json:"login_type"`
	Code      int    `json:"-"`

	Mobile string `json:"mobile,omitempty"`
}
