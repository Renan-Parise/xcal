package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renan-parise/gofreela/internal/entities"
	"github.com/renan-parise/gofreela/internal/positions"
	"github.com/renan-parise/gofreela/internal/repositories"
)

type createJobsRequest struct {
	Hash      string             `bson:"hash" json:"hash"`
	Positions positions.Position `bson:"positions" json:"positions"`
}

func CreateJobs(ctx *gin.Context, repo repositories.IJobsRepository) {
	var req createJobsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobs, err := entities.NewJobs(req.Hash, req.Positions)
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
