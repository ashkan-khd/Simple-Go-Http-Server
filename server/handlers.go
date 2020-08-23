package server

import (
	dbutils "Go-Server/Model/utils"
	"Go-Server/server/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func (server *Server) handleHelloWorld(responseWriter http.ResponseWriter, request *http.Request) {
	conn := server.newConnection(utils.Get, "/helloworld")
	resp := conn.Response
	resp.Message = "Hello world"
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(responseWriter).Encode(resp); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Oops Something Went Wrong!")
	}
}

func (server *Server) handleGetAllReqs() func(responseWriter http.ResponseWriter, request *http.Request) {
	if server.isHelpMode {
		return server.handleGetAllReqsAllowed
	} else {
		return server.NotAllowed
	}
}

func (server *Server) handleGetAllReqsAllowed(responseWriter http.ResponseWriter, request *http.Request) {
	resp := server.connections
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(responseWriter).Encode(resp); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Oops Something Went Wrong!")
	}
}

func (server *Server) NotAllowed(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	response := "Not Allowed To Access This Page :)"
	_, err := responseWriter.Write([]byte(response))
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Oops Something Went Wrong!")
	}
}

func (server *Server) handleGoodBye(responseWriter http.ResponseWriter, request *http.Request) {
	conn := server.newConnection(utils.Get, "/goodbye")
	responseWriter.WriteHeader(http.StatusNotFound)
	response := conn.Response
	response.Message = "status: 404"
	_, err := responseWriter.Write([]byte(response.Message))
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Oops Something Went Wrong!")
	}
}

func (server *Server) handleSignUp(responseWriter http.ResponseWriter, request *http.Request) {
	conn := server.newConnection(utils.Post, "/signup")
	var signedUpUser dbutils.User
	if err := json.NewDecoder(request.Body).Decode(&signedUpUser); err != nil {
		fmt.Println("Oops Something Went Wrong!")
	}
	byteArr, _ := json.Marshal(signedUpUser)
	conn.Request.Body = string(byteArr)
	responseWriter.Header().Set("Content-Type", "application/json")
	if server.jdb.ContainsUser(signedUpUser.Username) && server.jdb.GetUserByUsername(signedUpUser.Username).Password != signedUpUser.Password {
		conn.Response.Message = "Wrong Password"
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		if err := json.NewEncoder(responseWriter).Encode(conn.Response); err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Oops Something Went Wrong!")
		}
	} else {
		conn.AuthResponse.Message, conn.AuthResponse.Token = server.signUpAllowed(signedUpUser)
		responseWriter.WriteHeader(http.StatusAccepted)
		if err := json.NewEncoder(responseWriter).Encode(conn.AuthResponse); err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Oops Something Went Wrong!")
		}
	}
}

func (server *Server) signUpAllowed(signedUpUser dbutils.User) (message string, auth string) {
	if !server.jdb.ContainsUser(signedUpUser.Username) {
		message = "User " + signedUpUser.Username + " signed up successfully."
		server.jdb.PostUser(signedUpUser)
	} else {
		message = "User " + signedUpUser.Username + " logged In successfully."
	}
	auth = server.getAuthToken()
	server.jdb.PostAuthToken(auth)
	return
}

func (server *Server) getAuthToken() string {

	var auth [15]byte
	for i := 0; i < 15; i++ {
		var bt byte
		switch rand.Intn(4) {
		case 0:
			bt = 'A' + byte(rand.Intn(26))
		case 1:
			bt = 'a' + byte(rand.Intn(26))
		case 2:
			bt = '0' + byte(rand.Intn(10))
		case 3:
			bt = '#' + byte(rand.Intn(4))
		}
		auth[i] = bt
	}

	if server.jdb.ContainAuth(string(auth[:])) {
		return server.getAuthToken()
	}
	return string(auth[:])
}

func (server *Server) handleGetDb() func(responseWriter http.ResponseWriter, request *http.Request) {
	if server.isHelpMode {
		return server.handleGetDbAllowed
	} else {
		return server.NotAllowed
	}
}

func (server *Server) handleGetDbAllowed(responseWriter http.ResponseWriter, request *http.Request) {
	resp := server.jdb
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(responseWriter).Encode(resp); err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Oops Something Went Wrong!")
	}
}

func (server *Server) handleBuyProduct(responseWriter http.ResponseWriter, request *http.Request) {

}
