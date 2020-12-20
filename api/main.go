package main

import (
	"api/routes"
)

func main() {
	r := routes.SetRoutes()

	r.Run(":8081")
}