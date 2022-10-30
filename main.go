package main

import (
	"kanban-kita-api/config"
	"kanban-kita-api/routers"
	"log"
)

func main() {
	config.InitDB()
	r := routers.RoutesList()
	log.Fatal(r.Run(":3030"))
}
