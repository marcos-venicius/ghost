package main

import (
	"github.com/ghost"
)

func createProductHandler() string {
	return "create product handler"
}

func getProductHandler() string {
	return "get product handler"
}

func editProductHandler() string {
	return "edit product handler"
}

func main() {
	router := ghost.CreateRouter().Post(
		"/products", createProductHandler,
	).Get(
		"/products/:id", getProductHandler,
	).Put(
		"/products/:id", editProductHandler,
	)

	ghost.Listen(3030, router)
}
