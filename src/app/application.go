package app

import (
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
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	atHandler := http.NewHandler(
		accesstoken.NewService(
			db.NewRepository(),
		),
	)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8082")
}
