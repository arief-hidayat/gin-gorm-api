package route

import (
	"github.com/arief-hidayat/gin-gorm-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRoute(writerDb *gorm.DB, readerDb *gorm.DB, router *gin.Engine, logger *logger.Logger) {
	router.Static("/media", "/media")
	apiRouter := router.Group("/api/v1")
	contactRouter := apiRouter.Group("/contacts")
	ContactRoute(writerDb, readerDb, contactRouter, logger)
	//commentRouter := apiRouter.Group("/posts/:postId/comments")
	//CommentRoute(db, commentRouter, logger)
}
