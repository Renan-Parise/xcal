package router

import (
	"github.com/gin-gonic/gin"
	"github.com/renan-parise/gofreela/controllers"
	"github.com/renan-parise/gofreela/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

func JobsRouter(r *gin.Engine, db *mongo.Database) {
	v1 := r.Group("/v1/jobs")

	repo := repositories.NewJobsRepository(db)

	v1.POST("/", func(ctx *gin.Context) {
		controllers.CreateJobs(ctx, repo)
	})

	v1.GET("/:hash", func(ctx *gin.Context) {
		controllers.GetJobs(ctx, repo)
	})
}
