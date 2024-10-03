package database

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticConnection(cfg elasticsearch.Config) (*elasticsearch.Client, error) {
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	inf, err := es.Info()
	if err != nil {
		return nil, err
	}
	log.Println(inf)
	return es, nil
}
