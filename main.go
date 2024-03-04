package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrgmonsalve/back-event-booking/db"
	"github.com/jrgmonsalve/back-event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
