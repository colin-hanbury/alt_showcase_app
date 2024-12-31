package server

import (
	"log"
	"showcase-alt/server/src/actions"
	"showcase-alt/server/src/database"
	"showcase-alt/server/src/user"
	"showcase-alt/server/src/welcome"

	"github.com/gin-gonic/gin"
)

func Initialize() {

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
