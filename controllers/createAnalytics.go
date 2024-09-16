package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renan-parise/gofreela/internal/analytes"
	"github.com/renan-parise/gofreela/internal/entities"
	"github.com/renan-parise/gofreela/internal/repositories"
	"github.com/renan-parise/gofreela/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type createAnalyticsRequest struct {
	Hash     string            `bson:"hash" json:"hash"`
	Analytes analytes.Analytes `bson:"analytes" json:"analytes"`
}

func CreateAnalytics(ctx *gin.Context, repo repositories.IAnalyticsRepository) {
	var req createAnalyticsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	analytics, err := repo.Get(req.Hash)
	if err != nil && err != mongo.ErrNoDocuments {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if analytics != nil {
		analytics.Analytes.Merge(req.Analytes)
		analytics.UpdatedAt = *utils.Now()

		err = repo.Update(analytics)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		analytics, err = entities.NewAnalytics(req.Hash, req.Analytes)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = repo.Insert(analytics)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"success": true})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
