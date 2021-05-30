package app

import (
	"github.com/TestardR/bookstore_oauth-api/src/domain/access_token"
	"github.com/TestardR/bookstore_oauth-api/src/http"
	"github.com/TestardR/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartAppication() {
	atRepository := db.NewRepository()
	atService := access_token.NewService(atRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
