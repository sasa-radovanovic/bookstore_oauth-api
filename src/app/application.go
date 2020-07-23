package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/clients/cassandra"
	accesstoken "github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/access_token"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/http"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

// StartApplication starts an app
func StartApplication() {
	session := cassandra.GetSession()
	fmt.Println("Session fetched", !session.Closed())

	atHandler := http.NewHandler(
		accesstoken.NewService(
			db.NewRepository(),
		),
	)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8082")
}
