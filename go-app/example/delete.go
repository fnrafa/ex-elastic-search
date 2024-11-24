package example

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"io"
	"log"
)

func DeleteUser(es *elasticsearch.Client, name string) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": name,
			},
		},
	}
	jsonQuery, _ := json.Marshal(query)

	userResponse, err := es.Search(
		es.Search.WithIndex("users"),
		es.Search.WithBody(bytes.NewReader(jsonQuery)),
		es.Search.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error searching user: %s", err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(userResponse.Body)

	var userResult map[string]interface{}
	if err := json.NewDecoder(userResponse.Body).Decode(&userResult); err != nil {
		log.Fatalf("Error decoding user response: %s", err)
	}

	userHits := userResult["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(userHits) == 0 {
		fmt.Printf("User \"%s\" not found.\n", name)
		return
	}

	userDocID := userHits[0].(map[string]interface{})["_id"].(string)

	_, err = es.Delete(
		"users",
		userDocID,
		es.Delete.WithContext(context.Background()),
		es.Delete.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error deleting user: %s", err)
	}

	friendshipResponse, err := es.Search(
		es.Search.WithIndex("friendships"),
		es.Search.WithBody(bytes.NewReader(jsonQuery)),
		es.Search.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error searching friendship: %s", err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(friendshipResponse.Body)

	var friendshipResult map[string]interface{}
	if err := json.NewDecoder(friendshipResponse.Body).Decode(&friendshipResult); err != nil {
		log.Fatalf("Error decoding friendship response: %s", err)
	}

	friendshipHits := friendshipResult["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(friendshipHits) > 0 {
		friendshipDocID := friendshipHits[0].(map[string]interface{})["_id"].(string)
		_, err = es.Delete(
			"friendships",
			friendshipDocID,
			es.Delete.WithContext(context.Background()),
			es.Delete.WithRefresh("true"),
		)
		if err != nil {
			log.Fatalf("Error deleting friendship: %s", err)
		}
	}

	fmt.Printf("User \"%s\" and their friendship removed successfully.\n", name)
}
