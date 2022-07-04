package main

import "shop-pet/cmd"

// @title shop Swagger API
// @version 1.0
// @description this is document of shop service
// @termsOfService http://swagger.io/terms/
// @contact.name Amirali Kalhor
// @contact.url http://www.swagger.io/support
// @host localhost:8081
// @BasePath /
// @schemes http
func main() {
	cmd.Boot()

}
