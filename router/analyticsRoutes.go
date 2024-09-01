package router

import (
	"github.com/gin-gonic/gin"
	"github.com/renan-parise/xcal-analytics/controllers"
	"github.com/renan-parise/xcal-analytics/internal/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

func AnalyticsRouter(r *gin.Engine, db *mongo.Database) {
	v1 := r.Group("/v1/analytics")

	repo := repositories.NewAnalyticsRepository(db)

	v1.POST("/", func(ctx *gin.Context) {
		controllers.CreateAnalytics(ctx, repo)
	})

	v1.GET("/:hash", func(ctx *gin.Context) {
		controllers.GetAnalytics(ctx, repo)
	})
}
