package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gushikem01/go-handson/internals/usecase"
)

type UserHandler interface {
	AddRoutes(e *gin.Engine) *gin.Engine
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (uh *userHandler) AddRoutes(e *gin.Engine) *gin.Engine {
	g := e.Group("/api/v1/users")
	g.GET("/:id", uh.getUser)
	g.POST("", uh.createUser)
	g.PUT("/:id", uh.updateUser)
	g.DELETE("/:id", uh.deleteUser)
	return e
}

// getUser godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept  application/json
// @Produce  application/json
// @Param id path string true "ID"
// @Success 200 {object} UserResponse
// @Router /v1/users/{id} [get]
func (uh *userHandler) getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.userUsecase.FindUserById(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// createUser godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept  application/json
// @Produce  application/json
// @Param user body UserRequest true "User"
// @Success 200 {object} UserResponse
// @Router /v1/users [post]
func (uh *userHandler) createUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := uh.userUsecase.CreateUser(c, req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

// updateUser godoc
// @Summary Update user
// @Description Update user
// @Tags users
// @Accept  application/json
// @Produce  application/json
// @Param id path string true "ID"
// @Param user body UserRequest true "User"
// @Success 200 {object} UserResponse
// @Router /v1/users/{id} [put]
func (uh *userHandler) updateUser(c *gin.Context) {
	id := c.Param("id")
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := uh.userUsecase.UpdateUserById(c, id, req.Email, req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

// deleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags users
// @Accept  application/json
// @Produce  application/json
// @Param id path string true "ID"
// @Success 200
// @Router /v1/users/{id} [delete]
func (uh *userHandler) deleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := uh.userUsecase.DeleteUserById(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
