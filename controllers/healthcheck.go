package controllers

import (
	"github.com/gin-gonic/gin"
)

// HealthCheckController ....
type HealthCheckController struct{}

// HealthCheckSuccessResponse ...
type HealthCheckSuccessResponse struct {
	Message string `json:"message" binding:"required"`
}

// Index action: Get /healthehck
// @description ヘルスチェック用API
// @router /healthcheck [get]
// @Success 200 {object} HealthCheckSuccessResponse
// @tags healthcheck
func (pc HealthCheckController) Index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "success"})
	return
}
