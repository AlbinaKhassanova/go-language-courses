package main

import (
	"github.com/gin-gonic/gin"
	"go-language-courses/internal/database"
	"go-language-courses/internal/routes"
)

func main() {
	// Инициализация базы данных
	database.InitDatabase()

	// Создание маршрутов
	r := gin.Default()
	routes.SetupRoutes(r)

	// Запуск сервера
	r.Run(":8088")
}
