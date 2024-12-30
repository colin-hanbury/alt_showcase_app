package welcome

import (
	"context"
	"log"
	"net/http"
	"showcase-alt/server/src/shared/user"
	"time"

	"github.com/gin-gonic/gin"
)

type WelcomeController interface {
	GetWelcomeMessage(*gin.Context)
}

type welcomeController struct {
	userService user.UserService
}

func NewWelcomeController(userService user.UserService) WelcomeController {
	return &welcomeController{userService: userService}
}

func (controller *welcomeController) GetWelcomeMessage(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(3000)*time.Second)
	defer ctxErr()
	id := c.Param("id")
	message, err := controller.userService.GetWelcomeMessage(id, ctx)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": message})
}
