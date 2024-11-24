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

func GetUserDetails(es *elasticsearch.Client, name string) {
	userQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": name,
			},
		},
	}
	jsonUserQuery, _ := json.Marshal(userQuery)

	userResponse, err := es.Search(
		es.Search.WithIndex("users"),
		es.Search.WithBody(bytes.NewReader(jsonUserQuery)),
		es.Search.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error fetching user details: %s", err)
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

	user := userHits[0].(map[string]interface{})["_source"]
	fmt.Printf("Details of user \"%s\": %v\n", name, user)

	friendshipQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"user": name,
			},
		},
	}
	jsonFriendshipQuery, _ := json.Marshal(friendshipQuery)

	friendshipResponse, err := es.Search(
		es.Search.WithIndex("friendships"),
		es.Search.WithBody(bytes.NewReader(jsonFriendshipQuery)),
		es.Search.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error fetching friendship details: %s", err)
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
		friendship := friendshipHits[0].(map[string]interface{})["_source"]
		fmt.Printf("Friendship details for \"%s\": %v\n", name, friendship)
	} else {
		fmt.Printf("No friendship data found for \"%s\".\n", name)
	}
}
