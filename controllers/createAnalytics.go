package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renan-parise/xcal-analytics/internal/analytes"
	"github.com/renan-parise/xcal-analytics/internal/cases"
	"github.com/renan-parise/xcal-analytics/internal/repositories"
)

type createAnalyticsRequest struct {
	User     string            `bson:"user" json:"user"`
	Hash     string            `bson:"hash" json:"hash"`
	Analytes analytes.Analytes `bson:"analytes" json:"analytes"`
}

func CreateAnalytics(ctx *gin.Context, repo repositories.IAnalyticsRepository) {
	var req createAnalyticsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c := cases.NewCreateAnalyticsCase(repo)

	err := c.Execute(req.User, req.Hash, req.Analytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
