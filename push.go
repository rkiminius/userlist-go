package userlist

var (
	PUSH_KEY      string = ""
	PUSH_ID       string = ""
	PUSH_ENDPOINT string = "https://push.userlist.com"
)

type Config struct {
	PushKey      string `yaml:"pushKey"`
	PushID       string `yaml:"pushId"`
	PushEndpoint string `yaml:"endpoint"`
}

var Conf *Config

type UserList struct {
	Conf   *Config
	Client Client
	Users  User
	Events Event
}

func NewUserList() UserList {
	userlist := UserList{}
	userlist.Users.Scope = &userlist
	userlist.Events.Scope = &userlist
	return userlist
}
