package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"log"
)

func CheckAndDeleteIndex(ctx context.Context, es *elasticsearch.Client, index string) {
	res, err := es.Indices.Exists([]string{index}, es.Indices.Exists.WithContext(ctx))
	if err != nil {
		log.Fatalf("Error checking index existence: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode == 200 {
		fmt.Printf("Index \"%s\" exists. Deleting...\n", index)
		delRes, err := es.Indices.Delete([]string{index}, es.Indices.Delete.WithContext(ctx))
		if err != nil {
			log.Fatalf("Error deleting index \"%s\": %s", index, err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(delRes.Body)
	}
}

func CreateIndex(ctx context.Context, es *elasticsearch.Client, index string, mapping string) {
	_, err := es.Indices.Create(index, es.Indices.Create.WithContext(ctx), es.Indices.Create.WithBody(
		bytes.NewReader([]byte(mapping)),
	))
	if err != nil {
		log.Fatalf("Error creating index \"%s\": %s", index, err)
	}
}

func IndexData(ctx context.Context, es *elasticsearch.Client, index string, data []map[string]interface{}) {
	for _, item := range data {
		body, _ := json.Marshal(item)
		req := esapi.IndexRequest{
			Index:   index,
			Body:    bytes.NewReader(body),
			Refresh: "true",
		}
		res, err := req.Do(ctx, es)
		if err != nil {
			log.Fatalf("Error indexing data in \"%s\": %s", index, err)
		}
		err = res.Body.Close()
		if err != nil {
			return
		}
	}
}
