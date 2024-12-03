package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"project/class/database"
	_ "project/class/docs"
	"project/class/domain"
	"project/class/service"
)

type AuthController struct {
	service service.AuthService
	logger  *zap.Logger
	cacher  database.Cacher
}

func NewAuthController(service service.AuthService, logger *zap.Logger, rdb database.Cacher) *AuthController {
	return &AuthController{service: service, logger: logger, cacher: rdb}
}

// Login endpoint
// @Summary Admin login
// @Description authenticate user
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param domain.User body domain.User true " "
// @Success 200 {object} domain.HTTPResponse "Successful login"
// @Failure 404 {object} domain.HTTPResponse "User not found"
// @Failure 500 {object} domain.HTTPResponse "Internal server error"
// @Router  /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		responseError(c, "BIND_ERROR", err.Error(), http.StatusBadRequest)
		return
	}

	isAuthenticated, err := ctrl.service.Login(user)
	if err != nil {
		BadResponse(c, "server error", http.StatusInternalServerError)
		return
	}

	if !isAuthenticated {
		BadResponse(c, "authentication failed", http.StatusUnauthorized)
		return
	}

	token := uuid.New().String()
	IDKEY := user.Username

	if err := ctrl.cacher.Set(IDKEY, token); err != nil {
		BadResponse(c, "server error", http.StatusInternalServerError)
	}

	SuccessResponseWithData(c, "login successfully", http.StatusOK, gin.H{"token": token, "id_key": IDKEY})
}
