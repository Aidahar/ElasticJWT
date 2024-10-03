package elastic

import (
	"context"
	"encoding/json"
	"jwt/internal/domain"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/golang-jwt/jwt/v5"
)

type Places struct {
	es *elasticsearch.Client
}

func NewPlaces(es *elasticsearch.Client) *Places {
	return &Places{
		es: es,
	}
}

func (p *Places) GetPlaces(limit, offset int) (domain.Answer, int, error) {
	var places domain.Answer
	if offset > 1 {
		offset *= 10
	}
	res, err := p.es.Search(
		p.es.Search.WithContext(context.Background()),
		p.es.Search.WithIndex("magazins"),
		p.es.Search.WithFrom(offset-1),
		p.es.Search.WithSize(10),
		p.es.Search.WithSort("Location: asc"),
		p.es.Search.WithTrackTotalHits(true),
		p.es.Search.WithPretty(),
	)
	if err != nil {
		return places, 0, err
	}
	defer res.Body.Close()
	var mapResp map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&mapResp); err != nil {
		return places, 0, err
	}
	total := int(mapResp["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})
		source := doc["_source"]
		data, err := json.Marshal(source)
		if err != nil {
			return places, 0, err
		}
		var store domain.Store
		json.Unmarshal(data, &store)
		places.Places = append(places.Places, store)
	}
	places.Index = "Places"
	places.Total = total
	places.Last_page = total / 10
	return places, total, nil
}

func (p *Places) GetToken() domain.Token {
	var t domain.Token
	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2024, 4, 3, 20, 40, 00, 0, time.UTC).Unix(),
	})
	t.TokenStr, err = token.SignedString([]byte("HOLALAL"))
	if err != nil {
		log.Fatal(err)
	}
	return t
}
