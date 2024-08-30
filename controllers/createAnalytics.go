package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renan-parise/xcal-analytics/internal/cases"
	"github.com/renan-parise/xcal-analytics/internal/repositories"
)

type createAnalyticsRequest struct {
	Name        string `json:"name" binding:"required"`
	Information string `json:"information" binding:"required"`
	Values      []byte `json:"values" binding:"required"`
}

func CreateAnalytics(ctx *gin.Context, repo repositories.IAnalyticsRepository) {
	var req createAnalyticsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c := cases.NewCreateAnalyticsCase(repo)

	err := c.Execute(req.Name, req.Information, req.Values)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
