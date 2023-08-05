package route

import (
	"github.com/arief-hidayat/gin-gorm-api/internal/controller"
	"github.com/arief-hidayat/gin-gorm-api/internal/repository"
	"github.com/arief-hidayat/gin-gorm-api/internal/service"
	"github.com/arief-hidayat/gin-gorm-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ContactRoute(writerDb *gorm.DB, readerDb *gorm.DB, contactRouter *gin.RouterGroup, logger *logger.Logger) {
	var (
		contactRepository repository.ContactRepo = repository.NewContactRepo(writerDb, readerDb)
		contactService    service.ContactService = service.
					NewContactService(contactRepository)
		contactController controller.ContactController = controller.
					NewContactController(contactService, logger)
	)
	contactRouter.GET("", contactController.All)
	contactRouter.POST("", contactController.Insert)
	contactRouter.PUT(
		"/:contactId", contactController.Update)
	contactRouter.DELETE(
		"/:contactId", contactController.DeleteById)
	//contactRouter.POST("", middleware.AuthorizeJWT(), contactController.Insert)
	//contactRouter.PUT(
	//	"/:contactId", middleware.AuthorizeJWT(), contactController.Update)
	//contactRouter.DELETE(
	//	"/:contactId", middleware.AuthorizeJWT(), contactController.DeleteById)
}
