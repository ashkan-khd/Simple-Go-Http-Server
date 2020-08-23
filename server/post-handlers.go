package server

import (
	dbutils "Go-Server/model/utils"
	"Go-Server/server/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

func (server *Server) handleSignUp(responseWriter http.ResponseWriter, request *http.Request) {
	conn := server.newConnection(utils.Post, "/signup")
	signedUpUser := dbutils.User{Username: request.FormValue("username"), Password: request.FormValue("password")}
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

func (server *Server) handleBuyProduct(responseWriter http.ResponseWriter, request *http.Request) {
	conn := server.newConnection(utils.Post, "/buy_product")
	token := request.Header.Get("token")
	if server.jdb.ContainAuth(token) {
		product := dbutils.Product{Id: request.FormValue("product_id")}
		jsonBytes, _ := json.Marshal(product)
		conn.Request.Body = string(jsonBytes)
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conn.Response.Message = "product with id " + product.Id + " purchased successfully."
		if err := json.NewEncoder(responseWriter).Encode(conn.Response); err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Oops Something Went Wrong!")
		}
	} else {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		conn.Response.Message = "Wrong Auth-Token"
		if err := json.NewEncoder(responseWriter).Encode(conn.Response); err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Oops Something Went Wrong!")
		}
	}
}
