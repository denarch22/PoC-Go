package main

import (
	"sistema-usuarios-go/config"
	"sistema-usuarios-go/routes"
)

func main() {
	db := config.ConnectDatabase()
	r := routes.SetupRouter(db)
	r.Run(":9090")
}
