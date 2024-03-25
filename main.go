package main

import "final-project/handler"

func main() {

	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	// @description Type "Bearer" followed by a space and token you got from the User Login api.
	handler.StartApp()
}
