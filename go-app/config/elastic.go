package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func GetElasticClient() *elasticsearch.Client {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Username:  "elastic",
		Password:  "04tlHo4AJhYFhsNjC5e+",
	})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}
	return es
}
