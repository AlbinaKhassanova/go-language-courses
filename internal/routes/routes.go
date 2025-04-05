package routes

import (
	"github.com/gin-gonic/gin"
	"go-language-courses/internal/delivery"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/courses", delivery.CreateCourse)
	r.GET("/courses", delivery.GetCourses)
	r.GET("/courses/:id", delivery.GetCourseByID)
	r.PUT("/courses/:id", delivery.UpdateCourse)
	r.DELETE("/courses/:id", delivery.DeleteCourse)
}
