package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renan-parise/gofreela/internal/repositories"
)

func GetJobs(ctx *gin.Context, repo repositories.IJobsRepository) {
	hash := ctx.Param("hash")

	jobs, err := repo.Get(hash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if jobs == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Jobs not found"})
		return
	}

	ctx.JSON(http.StatusOK, jobs)
}
