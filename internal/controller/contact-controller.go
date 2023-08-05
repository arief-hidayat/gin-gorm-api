package controller

import (
	"fmt"
	"github.com/arief-hidayat/gin-gorm-api/internal/dto"
	"github.com/arief-hidayat/gin-gorm-api/internal/serializer"
	"github.com/arief-hidayat/gin-gorm-api/internal/service"
	"github.com/arief-hidayat/gin-gorm-api/internal/utils"
	"github.com/arief-hidayat/gin-gorm-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ContactController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	DeleteById(context *gin.Context)
}

type contactController struct {
	contactService service.ContactService
	logger         *logger.Logger
}

func NewContactController(
	contactService service.ContactService, logger *logger.Logger,
) *contactController {
	return &contactController{contactService: contactService, logger: logger}
}

// GetContacts             godoc
// @Summary      Get contacts list
// @Description  Responds with the list of all contacts as JSON.
// @Tags         contacts
// @Produce      json
// @Success      200  {object}  serializer.ContactResponse
// @Router       /contacts [get]
func (controller *contactController) All(context *gin.Context) {
	contacts := controller.contactService.All()
	serializer := serializer.ContactsSerializer{Contacts: contacts}
	context.JSON(http.StatusOK, utils.GetResponse(serializer.Response()))
}

// InsertContact             godoc
// @Summary      Insert contact
// @Description  Responds with contact as JSON.
// @Tags         contacts
// @Produce      json
// @Param data body dto.Contact true "Contact dto"
// @Success      201  {object}  serializer.ContactResponse
// @Router       /contacts [post]
func (controller *contactController) Insert(context *gin.Context) {
	contactDto := dto.Contact{}
	err := context.ShouldBindJSON(&contactDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	controller.logger.Info().Msg(fmt.Sprintf("institution=%s", contactDto.Institution))
	contact := controller.contactService.Insert(contactDto)
	controller.logger.Info().Msg(fmt.Sprintf("institution....=%s", contact.Institution))
	serializer := serializer.ContactSerializer{Contact: contact}
	context.JSON(http.StatusOK, utils.GetResponse(serializer.Response()))
}

// UpdateContact             godoc
// @Summary      Update contact
// @Description  Responds with contact as JSON.
// @Tags         contacts
// @Produce      json
// @Param        id  path      uint  true  "update contact by id"
// @Param data body dto.Contact true "Contact dto"
// @Success      200  {object}  serializer.ContactResponse
// @Router       /contacts/{id} [put]
func (controller *contactController) Update(context *gin.Context) {
	contactDto := dto.Contact{}
	err := context.ShouldBindJSON(&contactDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	contactId, _ := strconv.ParseUint(context.Param("contactId"), 10, 64)
	contact, err := controller.contactService.Update(uint(contactId), contactDto)
	if err != nil {
		context.JSON(http.StatusNotFound, utils.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	serializer := serializer.ContactSerializer{Contact: contact}
	context.JSON(http.StatusOK, utils.GetResponse(serializer.Response()))
}

// DeleteContact             godoc
// @Summary      Delete contact
// @Description  Responds with contact as JSON.
// @Tags         contacts
// @Produce      json
// @Param        id  path      uint  true  "delete contact by id"
// @Success      204
// @Router       /contacts/{id} [delete]
func (controller *contactController) DeleteById(context *gin.Context) {
	contactId, _ := strconv.ParseUint(context.Param("contactId"), 10, 64)
	result := controller.contactService.DeleteById(uint(contactId))
	if result.Error != nil {
		context.JSON(
			http.StatusBadRequest, utils.GetErrorResponse(result.Error.Error()))
		controller.logger.Error().Err(result.Error).Msg("")
		return
	} else if result.RowsAffected < 1 {
		context.JSON(
			http.StatusNotFound, utils.GetErrorResponse("contact does not exists"))
		controller.logger.Error().Msg("contact does not exists")
		return
	}
	context.JSON(http.StatusNoContent, utils.GetResponse(gin.H{}))
}
