package server

import (
	"Go-Server/Model"
	"Go-Server/server/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	port        string
	isHelpMode  bool
	connections []*utils.Connection
	router      *mux.Router
	jdb         *Model.JsonDatabase
}

func (server *Server) Run(isHelpMode bool) {
	server.isHelpMode = isHelpMode
	server.port = ":8008"
	server.router = mux.NewRouter()
	server.jdb = Model.InitDatabase()
	server.run()
}

func (server *Server) run() {
	fmt.Println("Response Starting To Listen On Port" + server.port + " ...")
	server.router.HandleFunc("/helloworld", server.handleHelloWorld)
	server.router.HandleFunc("/goodbyeworld", server.handleGoodBye)
	server.router.HandleFunc("/signup", server.handleSignUp)
	server.router.HandleFunc("/buy_product", server.handleBuyProduct)

	handleGetDb := server.handleGetDb()
	server.router.HandleFunc("/getdb", handleGetDb)
	handleGetAllReqFunc := server.handleGetAllReqs()
	server.router.HandleFunc("/getallreqs", handleGetAllReqFunc)
	if err := http.ListenAndServe(server.port, server.router); err != nil {
		fmt.Println("Oops Response Could Not Start Working!")
	}
}

func (server *Server) newConnection(method string, url string) *utils.Connection {
	conn := utils.GetConnection(method, url)
	server.connections = append(server.connections, conn)
	return conn
}
