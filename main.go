package main

import (
	"go-jwt/configs"
	"go-jwt/routers"
)

func main() {
	configs.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
