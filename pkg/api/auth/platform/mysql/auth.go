package mysql

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"gorm.io/gorm"

	"github.com/zaynkorai/eventus"
)

// User represents the client for user table
type User struct{}

// Custom errors
var (
	ErrAlreadyExists   = echo.NewHTTPError(http.StatusConflict, "Username or email already exists.")
	ErrInsertFailed    = echo.NewHTTPError(http.StatusInternalServerError, "Unable to create user .")
	ErrNoUserWithEmail = echo.NewHTTPError(http.StatusBadRequest, "User not available with given email")
)

// Create creates a new user on database
func (u User) Create(db *gorm.DB, usr eventus.User) (eventus.User, error) {
	var user = new(eventus.User)

	result := db.First(&user, "lower(username) = ? or lower(email) = ?",
		strings.ToLower(usr.Username), strings.ToLower(usr.Email))

	if result.RowsAffected > 0 {
		return eventus.User{}, errors.New("Username or email already exists")
	}

	db.Create(&usr)

	return usr, nil
}

// Read returns single user by ID
func (u User) Read(db *gorm.DB, id int) (eventus.User, error) {
	var user eventus.User

	result := db.First(&user, "id = ?", id)

	if err := result.Error; err != nil {
		return eventus.User{}, errors.New("No User with that ID exists")
	}

	return user, nil
}

// FindByUsername queries for single user by username
func (u User) FindByUsername(db *gorm.DB, uname string) (eventus.User, error) {

	var user eventus.User
	result := db.First(&user, "username = ?", uname)

	if err := result.Error; err != nil {
		return eventus.User{}, errors.New("No User with that Username exists")
	}

	return user, nil
}

// FindByEmail queries for single user by email
func (u User) FindByEmail(db *gorm.DB, email string) (eventus.User, error) {
	var user eventus.User
	result := db.First(&user, "email = ?", email)

	if err := result.Error; err != nil {
		return eventus.User{}, errors.New("No User with that Email exists")
	}

	return user, nil
}

// FindByUser queries for single user by username
func (u User) FindByUser(db *gorm.DB, usr eventus.User) (eventus.User, error) {
	var user eventus.User
	result := db.First(&user, "username = ? and email = ?", usr.Username, usr.Email)

	if err := result.Error; err != nil {
		return eventus.User{}, errors.New("No User with that username and email exists")
	}

	return user, nil
}

// FindByToken queries for single user by token
func (u User) FindByToken(db *gorm.DB, token string) (eventus.User, error) {
	var user eventus.User

	result := db.First(&user, "token = ?", token)
	if err := result.Error; err != nil {
		return eventus.User{}, errors.New("No User with that Email exists")
	}

	return user, nil
}

// Update updates user's info
func (u User) Update(db *gorm.DB, user eventus.User) error {
	db.Model(&user).Updates(user)
	return nil
}
