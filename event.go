package userlist

type Event struct {
	Scope      *UserList
	Name       string      `json:"name"`
	User       string      `json:"user"`
	Properties interface{} `json:"properties,omitempty"`
}

func (e Event) endpoint() string {
	return "/events"
}

func (e Event) Push(event Event) {
	e.Scope.Client.Post(e.endpoint(), event)
}
