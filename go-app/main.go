package main

import "go-app/initializer"

func main() {
	initializer.Elasticsearch()

	/*es := config.GetElasticClient()*/

	// Step 1: Create a new user and their friendships
	/*example.CreateNewUser(es, types.User{
		Name:    "Frank",
		Age:     29,
		Hobbies: []string{"gaming", "traveling"},
	}, []string{"Eve", "Dave"})*/

	// Step 2: Get user details and their friendships
	/*example.GetUserDetails(es, "Frank")*/

	// Step 3: Update user details and friendships
	/*example.UpdateUser(es, "Frank", types.User{
		Name:    "Frank",
		Age:     30,
		Hobbies: []string{"hiking", "coding"},
	}, []string{"Alice", "Bob"})*/

	// Step 4: Get updated user details
	/*example.GetUserDetails(es, "Frank")*/

	// Step 5: Delete the user and their friendships
	/*example.DeleteUser(es, "Frank")*/

	// Step 6: Attempt to get details of the deleted user
	/*example.GetUserDetails(es, "Frank")*/
}
