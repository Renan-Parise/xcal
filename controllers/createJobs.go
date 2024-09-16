package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renan-parise/gofreela/internal/entities"
	"github.com/renan-parise/gofreela/internal/positions"
	"github.com/renan-parise/gofreela/internal/repositories"
	"github.com/renan-parise/gofreela/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type createJobsRequest struct {
	Hash      string              `bson:"hash" json:"hash"`
	Positions positions.Positions `bson:"positions" json:"positions"`
}

func CreateJobs(ctx *gin.Context, repo repositories.IJobsRepository) {
	var req createJobsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobs, err := repo.Get(req.Hash)
	if err != nil && err != mongo.ErrNoDocuments {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if jobs != nil {
		jobs.Positions.Merge(req.Positions)
		jobs.UpdatedAt = *utils.Now()

		err = repo.Update(jobs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		jobs, err = entities.NewJobs(req.Hash, req.Positions)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = repo.Insert(jobs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"success": true})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
