package rest

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}
	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}
	return jwtToken[1], nil
}

func jwtMiddleware(c *gin.Context, tokenStr string) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	if jwtToken != tokenStr {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	} else {
		c.JSON(http.StatusOK, "HELLO BODY")
	}
	c.Next()
}
