package utils

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authtoken struct {
	Token string `json:"token"`
}

type Product struct {
	Id string `json:"product_id"`
}
