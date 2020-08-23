package utils

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

type Authtoken struct {
	Token string `json "token"`
}
