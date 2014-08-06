package main

import "github.com/chazzuka/go/martinier/martinier"

func main() {
	db := martinier.NewConnection("martinier", "192.168.10.10:27017")
	server := martinier.NewServer(db)
	server.Run()
}
