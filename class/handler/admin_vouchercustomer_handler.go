package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/class/service"
	"project/class/util"
)

type AdminVoucherCustomerController struct {
	service service.VoucherService
	logger  *zap.Logger
}

func NewAdminVoucherCustomerController(service service.VoucherService, logger *zap.Logger) *AdminVoucherCustomerController {
	return &AdminVoucherCustomerController{service, logger}
}

func (ctrl *AdminVoucherCustomerController) Get(c *gin.Context) {
	page, err := util.Uint(c.Query("p"))
	limit, err := util.Uint(c.Query("l"))
	isActive := c.Query("active")
	area := c.Query("area")
	voucherType := c.Query("type")

	total, pages, currentPage, itemsPerPage, vouchers, err := ctrl.service.All(page, limit, isActive, area, voucherType)
	if err != nil {
		ctrl.logger.Error("Failed to get vouchers", zap.Error(err))
		responseError(c, "GET_ERROR", "Failed to get vouchers", http.StatusInternalServerError)
		return
	}

	ctrl.logger.Info("Vouchers retrieved successfully", zap.Any("vouchers", vouchers))
	responseDataPage(c, total, pages, currentPage, itemsPerPage, vouchers)
}
