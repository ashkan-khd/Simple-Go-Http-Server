package utils

type response struct {
	Message string `json: "message"`
}

type authResponse struct {
	*response
	Token string `json: token`
}

type request struct {
	Body string `json: "body"`
}

type Connection struct {
	Method       string        `json: "mehtod"`
	URL          string        `json: "url"`
	Request      *request      `json: "client"`
	Response     *response     `json: "server-response"`
	AuthResponse *authResponse `json: "server-auth-response"`
}

func GetConnection(method string, url string) *Connection {
	conn := &Connection{Method: method}
	conn.URL = url
	conn.Request = &request{}
	conn.Response = &response{}
	conn.AuthResponse = &authResponse{response: &response{}}
	return conn
}

var Get string = "GET"
var Post string = "POST"
