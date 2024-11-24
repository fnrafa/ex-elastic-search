package initializer

import (
	"context"
	"fmt"
	"go-app/config"
	"go-app/types"
	"go-app/utils"
	"go-app/utils/elastic"
)

func Elasticsearch() {
	ctx := context.Background()
	es := config.GetElasticClient()

	users := utils.ReadJSON[[]types.User]("../shared/users.json")
	friendships := utils.ReadJSON[[]types.Friendship]("../shared/friendships.json")

	for _, index := range []string{"users", "friendships"} {
		elastic.CheckAndDeleteIndex(ctx, es, index)
	}

	fmt.Println("Creating \"users\" index...")
	userMapping := `{
		"mappings": {
			"properties": {
				"name": { "type": "text", "fields": { "keyword": { "type": "keyword" } } },
				"age": { "type": "integer" },
				"hobbies": { "type": "text", "fields": { "keyword": { "type": "keyword" } } }
			}
		}
	}`
	elastic.CreateIndex(ctx, es, "users", userMapping)

	fmt.Println("Creating \"friendships\" index...")
	friendshipMapping := `{
		"mappings": {
			"properties": {
				"user": { "type": "keyword" },
				"friends": { "type": "keyword" }
			}
		}
	}`
	elastic.CreateIndex(ctx, es, "friendships", friendshipMapping)

	fmt.Println("Indexing user data...")
	userData := make([]map[string]interface{}, len(users))
	for i, user := range users {
		userData[i] = map[string]interface{}{
			"name":    user.Name,
			"age":     user.Age,
			"hobbies": user.Hobbies,
		}
	}
	elastic.IndexData(ctx, es, "users", userData)

	fmt.Println("Indexing friendship data...")
	friendshipData := make([]map[string]interface{}, len(friendships))
	for i, friendship := range friendships {
		friendshipData[i] = map[string]interface{}{
			"user":    friendship.User,
			"friends": friendship.Friends,
		}
	}
	elastic.IndexData(ctx, es, "friendships", friendshipData)

	fmt.Println("Elasticsearch initialization complete.")
}
