package server

import (
	"log"
	"showcase-alt/server/src/shared/actions"
	"showcase-alt/server/src/shared/database"
	"showcase-alt/server/src/shared/user"
	"showcase-alt/server/src/shared/welcome"

	"github.com/gin-gonic/gin"
)

func initialize() {

	client, err := database.ConnectMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewUserRepo(client)
	userService := user.NewUserService(userRepo)
	welcomeController := welcome.NewWelcomeController(userService)
	actionsController := actions.NewActionsController(userService)

	router := gin.Default()

	router.POST("/actions", actionsController.AddUser)
	router.GET("/welcome", welcomeController.GetWelcomeMessage)

	router.Run("localhost:8080")
}
