package main

import (
	"github.com/nataliatuhten/loja-digiport-backend/routes"
)

func StartServer() {
	routes.HandleRequests()
}
