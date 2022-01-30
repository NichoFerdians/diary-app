package http

import (
	"diary-app-service/entity"
	"diary-app-service/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase usecase.UserUseCaseInterface
}

func NewUserHandler(useCase usecase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		useCase: useCase,
	}
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var input entity.LoginUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.useCase.LoginUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp := entity.LoginUserResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, gin.H{"data": rsp})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input entity.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.useCase.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp := entity.UserResponse{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, gin.H{"data": rsp})
}
