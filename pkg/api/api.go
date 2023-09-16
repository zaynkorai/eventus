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
	"os"

	"github.com/zaynkorai/eventus/pkg/utl/mysql"
	"github.com/zaynkorai/eventus/pkg/utl/server"

	"github.com/zaynkorai/eventus/pkg/api/event"
	eventTransport "github.com/zaynkorai/eventus/pkg/api/event/transport"
)

// Start starts the API service
func Start() error {
	dbURL := os.Getenv("DATABASE_URL")
	serverPort := ":" + os.Getenv("PORT")

	db, err := mysql.New(dbURL, 30)
	if err != nil {
		return err
	}

	e := server.New()
	v1 := e.Group("/v1")

	eventTransport.NewHTTP(event.Initialize(db), v1)

	server.Start(e, &server.Config{
		Port:                serverPort,
		ReadTimeoutSeconds:  30,
		WriteTimeoutSeconds: 10,
		Debug:               true,
	})

	return nil
}
