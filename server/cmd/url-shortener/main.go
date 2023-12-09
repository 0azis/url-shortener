package main

import (
	"fmt"
	"url-shortener/internal/pkg"
	"url-shortener/internal/server"
)

func main() {
	fmt.Println(pkg.GenerateUUID(6))
	server.InitServer()
}
