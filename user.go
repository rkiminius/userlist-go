package userlist

type User struct {
	Scope      *UserList
	Identifier string      `json:"identifier"`
	Email      string      `json:"email"`
	Properties interface{} `json:"properties,omitempty"`
}

func (u User) endpoint() string {
	return "/users"
}

func (u User) Push(user User) {
	u.Scope.Client.Post(u.endpoint(), user)
}

func (u User) Delete(identifier string) {
	endpoint := u.endpoint() + "/" + identifier
	u.Scope.Client.Delete(endpoint)
}
