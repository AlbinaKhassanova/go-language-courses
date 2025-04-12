package main

import (
	"github.com/gin-gonic/gin"
	"go-language-courses/internal/db"
	"go-language-courses/internal/routes"
)

func main() {

	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8088")
}
