package auth

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"

	"github.com/zaynkorai/eventus"
	"github.com/zaynkorai/eventus/pkg/api/auth/platform/mysql"
)

// New creates new iam service
func New(db *gorm.DB, udb UserDB, j TokenGenerator, sec Securer) Auth {
	return Auth{
		db:  db,
		udb: udb,
		tg:  j,
		sec: sec,
	}
}

// Initialize initializes auth application service
func Initialize(db *gorm.DB, j TokenGenerator, sec Securer) Auth {
	return New(db, mysql.User{}, j, sec)
}

// Service represents auth service interface
type Service interface {
	CreateUser(echo.Context, eventus.User) (eventus.User, error)
	Authenticate(echo.Context, string, string) (eventus.AuthToken, error)
	Refresh(echo.Context, string) (string, error)
}

// Auth represents auth application service
type Auth struct {
	db  *gorm.DB
	udb UserDB
	tg  TokenGenerator
	sec Securer
}

// UserDB represents user repository interface
type UserDB interface {
	Create(*gorm.DB, eventus.User) (eventus.User, error)
	Read(*gorm.DB, int) (eventus.User, error)
	FindByUsername(*gorm.DB, string) (eventus.User, error)
	FindByToken(*gorm.DB, string) (eventus.User, error)
	Update(*gorm.DB, eventus.User) error
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateToken(eventus.User) (string, error)
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
	HashMatchesPassword(string, string) bool
	Token(string) string
	ResetToken() string
}
