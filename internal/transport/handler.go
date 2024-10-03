package rest

import (
	"fmt"
	"jwt/internal/domain"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var tokenSecret = []byte("Secret token!")

var tokens []string

type Places interface {
	GetPlaces(limit, offset int) (domain.Answer, int, error)
	GetToken() domain.Token
}

type Handler struct {
	placesServ Places
}

func NewHandler(placesServ Places) *Handler {
	return &Handler{
		placesServ: placesServ,
	}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	r.GET("/api/places", h.GetPlaces)
	r.GET("/api/get_token", h.GetToken)
	r.GET("/api/recommend", h.Recomend)
}

func (h *Handler) Recomend(c *gin.Context) {
	if len(tokens) > 0 {
		jwtMiddleware(c, tokens[0])
	} else {
		c.JSON(http.StatusUnauthorized, "bad jwt token")
	}
}

func (h *Handler) GetToken(c *gin.Context) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(tokenSecret)
	c.Header("Content-Type", "application/json")
	tokens = append(tokens, tokenString)
	c.JSON(http.StatusOK, tokenString)
}

func (h *Handler) GetPlaces(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		e := fmt.Sprintf("error: Invalid page value %s", pageStr)
		c.JSON(http.StatusBadRequest, e)
	}
	places, total, err := h.placesServ.GetPlaces(20000, page)
	if err != nil {
		log.Fatal(err)
	}

	pageCount := total / 10
	if pageCount == 0 {
		pageCount = 1
	}
	if page < 1 || page > pageCount {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var prevPage, nextPage int
	if page > 1 {
		prevPage = page - 1
	}
	if page < pageCount {
		nextPage = page + 1
	}
	places.Prev_Page = prevPage
	places.Next_Page = nextPage
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, places)
}
