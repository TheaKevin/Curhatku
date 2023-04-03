package api

import (
	"Curhatku/backend/controllers/authentication"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "DELETE", "PUT", "GET", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	authRepo := authentication.NewRepository(s.DB)
	authService := authentication.NewService(authRepo)
	authHandler := authentication.NewHandler(authService)
	s.Router.POST("/login", authHandler.Login)
	s.Router.POST("/register", authHandler.Register)
	s.Router.GET("/user", authHandler.AuthenticateUser)
	s.Router.POST("/logout", authHandler.Logout)
	s.Router.PATCH("/changePassword", authHandler.ChangePassword)

}
