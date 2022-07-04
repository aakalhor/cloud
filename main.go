package main

import "order-pet/cmd"

// @title order Swagger API
// @version 1.0
// @description this is document of order service
// @termsOfService http://swagger.io/terms/
// @contact.name Amirali Kalhor
// @contact.url http://www.swagger.io/support
// @host localhost:8081
// @BasePath /
// @schemes http
func main() {
	cmd.Boot()
}
