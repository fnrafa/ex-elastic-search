package types

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

type Friendship struct {
	User    string   `json:"user"`
	Friends []string `json:"friends"`
}
