package app

import (
	"github.com/TestardR/bookstore_oauth-api/src/domain/access_token"
	"github.com/TestardR/bookstore_oauth-api/src/repository/db"
)

func StartAppication() {
	dbRepository := db.New()
	atService := access_token.NewService(dbRepository)
}
