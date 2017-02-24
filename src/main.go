package main

import (
	"server"
)

func main() {
	server:=&server.Server{Port:"8852"}
	server.Start()

}
