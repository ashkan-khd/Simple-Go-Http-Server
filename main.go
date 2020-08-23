package main

import "Go-Server/server"

func main() {
	ser := server.Server{}
	ser.Run(true)
}
