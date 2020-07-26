package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	accesstoken "github.com/sasa-radovanovic/bookstore_oauth-api/src/domain/access_token"
	"github.com/sasa-radovanovic/bookstore_oauth-api/src/utils/errors"
)

// AccessTokenHandler type
type AccessTokenHandler interface {
	GetByID(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

// NewHandler is the new handler
func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

// GetById is actual handler
func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := h.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

// Create is actual handler for creating new token
func (h *accessTokenHandler) Create(c *gin.Context) {
	var atRequest accesstoken.AtRequest
	if err := c.ShouldBindJSON(&atRequest); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}
	accessToken, err := h.service.Create(atRequest)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
