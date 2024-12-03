package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/class/service"
)

type CustomerRedeemVoucherController struct {
	service service.RedeemVoucherService
	logger  *zap.Logger
}

func NewCustomerRedeemVoucherController(service service.RedeemVoucherService, logger *zap.Logger) *CustomerRedeemVoucherController {
	return &CustomerRedeemVoucherController{service, logger}
}

func (ctrl *CustomerRedeemVoucherController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Redemption created"})
}
