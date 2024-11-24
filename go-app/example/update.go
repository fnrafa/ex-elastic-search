package example

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"go-app/types"
	"io"
	"log"
)

func UpdateUser(es *elasticsearch.Client, name string, updatedData types.User, updatedFriends []string) {
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"user": name,
			},
		},
	}
	jsonSearchQuery, _ := json.Marshal(searchQuery)

	friendshipResponse, err := es.Search(
		es.Search.WithIndex("friendships"),
		es.Search.WithBody(bytes.NewReader(jsonSearchQuery)),
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
	if len(friendshipHits) == 0 {
		fmt.Printf("Friendship for \"%s\" not found.\n", name)
		return
	}

	friendshipDocID := friendshipHits[0].(map[string]interface{})["_id"].(string)

	updateFriendship := map[string]interface{}{
		"doc": types.Friendship{
			User:    name,
			Friends: updatedFriends,
		},
	}
	jsonUpdateFriendship, _ := json.Marshal(updateFriendship)

	updatedFriendshipResponse, err := es.Update(
		"friendships",
		friendshipDocID,
		bytes.NewReader(jsonUpdateFriendship),
		es.Update.WithContext(context.Background()),
		es.Update.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error updating friendship: %s", err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(updatedFriendshipResponse.Body)

	fmt.Printf("Friendship for \"%s\" updated successfully.\n", name)

	userSearchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": name,
			},
		},
	}
	jsonUserSearchQuery, _ := json.Marshal(userSearchQuery)

	userSearchResponse, err := es.Search(
		es.Search.WithIndex("users"),
		es.Search.WithBody(bytes.NewReader(jsonUserSearchQuery)),
		es.Search.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error searching user: %s", err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(userSearchResponse.Body)

	var userResult map[string]interface{}
	if err := json.NewDecoder(userSearchResponse.Body).Decode(&userResult); err != nil {
		log.Fatalf("Error decoding user response: %s", err)
	}

	userHits := userResult["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(userHits) == 0 {
		fmt.Printf("User \"%s\" not found.\n", name)
		return
	}

	userDocID := userHits[0].(map[string]interface{})["_id"].(string)

	updateUser := map[string]interface{}{
		"doc": updatedData,
	}
	jsonUpdateUser, _ := json.Marshal(updateUser)

	updatedUserResponse, err := es.Update(
		"users",
		userDocID,
		bytes.NewReader(jsonUpdateUser),
		es.Update.WithContext(context.Background()),
		es.Update.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error updating user: %s", err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(updatedUserResponse.Body)

	fmt.Printf("User \"%s\" updated successfully.\n", name)
}
