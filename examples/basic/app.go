package main

import (
	ghost "github.com/marcos-venicius/ghost"
)

func createProductHandler() interface{} {
	return "create product handler"
}

func getProductHandler() interface{} {
	return "get product handler"
}

func editProductHandler() interface{} {
	return "edit product handler"
}

func main() {
	router := ghost.CreateRouter()

	router.Post(
		"/products", createProductHandler,
	)

	router.Get(
		"/products/:id", getProductHandler,
	)

	router.Put(
		"/products/:id", editProductHandler,
	)

	router.Listen(3030)
}
