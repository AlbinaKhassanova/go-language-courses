package routes

import (
	"github.com/gin-gonic/gin"
	"go-language-courses/internal/auth"
	"go-language-courses/internal/delivery"
	"go-language-courses/internal/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/courses", delivery.CreateCourse)
	r.GET("/courses", delivery.GetCourses)
	r.GET("/courses/:id", delivery.GetCourseByID)
	r.PUT("/courses/:id", delivery.UpdateCourse)
	r.DELETE("/courses/:id", delivery.DeleteCourse)

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	protected := r.Group("api")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me)
	}
}
