// Eventus - Go(lang) Restful API
//
// API Docs for Eventus v1
//
// 	 Terms Of Service:  N/A
//     Schemes: http
//     Version: 1.0.0
//     Contact: Zayn Korai <zaynulabdin313@gmail.com>
//     Host: localhost:8080
//
package api

import (
	"crypto/sha1"
	"os"

	"github.com/zaynkorai/eventus/pkg/utl/jwt"
	authMw "github.com/zaynkorai/eventus/pkg/utl/middleware/auth"
	"github.com/zaynkorai/eventus/pkg/utl/mysql"
	"github.com/zaynkorai/eventus/pkg/utl/secure"
	"github.com/zaynkorai/eventus/pkg/utl/server"
	"github.com/zaynkorai/eventus/pkg/utl/zlog"

	"github.com/zaynkorai/eventus/pkg/api/event"
	eventLogging "github.com/zaynkorai/eventus/pkg/api/event/logging"
	eventTransport "github.com/zaynkorai/eventus/pkg/api/event/transport"

	"github.com/zaynkorai/eventus/pkg/api/auth"
	authLogging "github.com/zaynkorai/eventus/pkg/api/auth/logging"
	authTransport "github.com/zaynkorai/eventus/pkg/api/auth/transport"
)

// Start starts the API service
func Start() error {
	dbURL := os.Getenv("DATABASE_URL")
	serverPort := ":" + os.Getenv("PORT")

	db, err := mysql.New(dbURL, 30)
	if err != nil {
		return err
	}

	log := zlog.New()
	jwt, err := jwt.New("HS256", os.Getenv("JWT_SECRET"), 15, 64)
	sec := secure.New(8, sha1.New(), os.Getenv("KEY_STRING"))
	if err != nil {
		return err
	}

	e := server.New()
	v1 := e.Group("/v1")
	authMiddleware := authMw.Middleware(jwt)

	eventTransport.NewHTTP(eventLogging.New(event.Initialize(db), log), v1)
	authTransport.NewHTTP(authLogging.New(auth.Initialize(db, jwt, sec), log), e, authMiddleware)

	server.Start(e, &server.Config{
		Port:                serverPort,
		ReadTimeoutSeconds:  30,
		WriteTimeoutSeconds: 10,
		Debug:               true,
	})

	return nil
}
