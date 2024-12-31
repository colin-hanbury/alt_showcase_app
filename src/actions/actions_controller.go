package actions

import (
	"context"
	"log"
	"net/http"
	"showcase-alt/server/src/shared/user"
	"time"

	"github.com/gin-gonic/gin"
)

type ActionsController interface {
	AddUser(*gin.Context)
}

type actionsController struct {
	userService user.UserService
}

func NewActionsController(userService user.UserService) ActionsController {
	return &actionsController{userService: userService}
}

func (controller *actionsController) AddUser(c *gin.Context) {
	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(3000)*time.Second)
	defer ctxErr()
	userInput := user.UserInput{Name: c.Param("name"), Nationality: c.Param("nationality")}
	user, err := controller.userService.AddUser(userInput, ctx)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}
