package app

import (
	"github.com/TestardR/bookstore_oauth-api/src/clients/cassandra"
	"github.com/TestardR/bookstore_oauth-api/src/domain/access_token"
	"github.com/TestardR/bookstore_oauth-api/src/http"
	"github.com/TestardR/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartAppication() {
	cassandra.GetSession()

	atRepository := db.NewRepository()
	atService := access_token.NewService(atRepository)
	atHandler := http.NewAccessTokenHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/:access_token_id", atHandler.Create)

	router.Run(":8080")
}
