package route

import (
	"github.com/arief-hidayat/gin-gorm-api/internal/utils"
	"github.com/arief-hidayat/gin-gorm-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type healthCheckController struct {
}

func NewHealthCheckController() *healthCheckController {
	return &healthCheckController{}
}

func (c *healthCheckController) Check(context *gin.Context) {
	context.JSON(http.StatusOK, utils.Response{Data: nil, Status: true, Errors: nil})
}

func RootRoute(writerDb *gorm.DB, readerDb *gorm.DB, router *gin.Engine, logger *logger.Logger) {
	router.Static("/media", "/media")
	router.GET("/healthz", NewHealthCheckController().Check)
	apiRouter := router.Group("/api/v1")
	contactRouter := apiRouter.Group("/contacts")
	ContactRoute(writerDb, readerDb, contactRouter, logger)
	//commentRouter := apiRouter.Group("/posts/:postId/comments")
	//CommentRoute(db, commentRouter, logger)
}
