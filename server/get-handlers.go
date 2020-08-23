package server

import (
	"Go-Server/server/utils"
	"encoding/json"
	"fmt"
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
