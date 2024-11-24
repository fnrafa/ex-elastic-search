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

func CreateNewUser(es *elasticsearch.Client, user types.User, friends []string) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshalling user data: %s", err)
	}

	response, err := es.Index(
		"users",
		bytes.NewReader(jsonData),
		es.Index.WithContext(context.Background()),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error creating new user: %s", err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(response.Body)

	fmt.Println("New user created:", response.Status())

	friendship := types.Friendship{
		User:    user.Name,
		Friends: friends,
	}
	jsonFriendship, err := json.Marshal(friendship)
	if err != nil {
		log.Fatalf("Error marshalling friendship data: %s", err)
	}

	friendshipResponse, err := es.Index(
		"friendships",
		bytes.NewReader(jsonFriendship),
		es.Index.WithContext(context.Background()),
		es.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Fatalf("Error creating friendship: %s", err)
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(friendshipResponse.Body)

	fmt.Println("Friendship created:", friendshipResponse.Status())
}
