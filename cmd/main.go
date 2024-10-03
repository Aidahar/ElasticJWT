package main

import (
	"jwt/internal/repository/elastic"
	"jwt/internal/service"
	rest "jwt/internal/transport"
	database "jwt/pkg/database"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "123456",
	}
	db, err := database.NewElasticConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	placesRepo := elastic.NewPlaces(db)
	placesServ := service.NewServices(placesRepo)
	handler := rest.NewHandler(placesServ)

	router := gin.Default()
	handler.InitRoutes(router)
	err = router.Run(":8888")
	if err != nil {
		log.Fatal(err)
	}
}
