package app

import (
	http "github.com/ayush723/oauth-api_bookstore/src/http/access_token"
	"github.com/gin-gonic/gin"

	"github.com/ayush723/oauth-api_bookstore/src/repository/db"

	"github.com/ayush723/oauth-api_bookstore/src/domain/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
