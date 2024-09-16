package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renan-parise/gofreela/internal/repositories"
)

func GetAnalytics(ctx *gin.Context, repo repositories.IAnalyticsRepository) {
	hash := ctx.Param("hash")

	analytics, err := repo.Get(hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if analytics == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Analytics not found"})
		return
	}

	ctx.JSON(http.StatusOK, analytics)
}
