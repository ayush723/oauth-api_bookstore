package app

import (
	"github.com/ayush723/oauth-api_bookstore/src/clients/cassandra"
	http "github.com/ayush723/oauth-api_bookstore/src/http/access_token"
	"github.com/ayush723/oauth-api_bookstore/src/repository/db"
	"github.com/ayush723/oauth-api_bookstore/src/repository/rest"
	"github.com/ayush723/oauth-api_bookstore/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	session := cassandra.GetSession()

	defer session.Close()
	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
