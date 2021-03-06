// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql_model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userID"`
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Todos []*Todo `json:"todos"`
}
