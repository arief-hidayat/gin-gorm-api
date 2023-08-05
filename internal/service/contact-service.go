package service

import (
	"github.com/arief-hidayat/gin-gorm-api/internal/dto"
	"github.com/arief-hidayat/gin-gorm-api/internal/model"
	"github.com/arief-hidayat/gin-gorm-api/internal/repository"
	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type ContactService interface {
	All() []model.Contact
	Insert(contactDto dto.Contact) model.Contact
	Update(contactId uint, contactDto dto.Contact) (model.Contact, error)
	DeleteById(contactId uint) *gorm.DB
}

type contactService struct {
	contactRepo repository.ContactRepo
}

func NewContactService(contactRepo repository.ContactRepo) *contactService {
	return &contactService{contactRepo: contactRepo}
}

func (service *contactService) All() []model.Contact {
	return service.contactRepo.AllContacts()
}

func (service *contactService) Insert(
	contactDto dto.Contact,
) model.Contact {
	contactModel := model.Contact{}
	err := smapping.FillStruct(&contactModel, smapping.MapFields(&contactDto))
	if err != nil {
		panic(err)
	}
	return service.contactRepo.Insert(contactModel)
}

func (service *contactService) Update(
	contactId uint, contactDto dto.Contact,
) (model.Contact, error) {
	contact, err := service.contactRepo.GetById(contactId)
	if err != nil {
		return contact, err
	}
	fillErr := smapping.FillStruct(&contact, smapping.MapFields(&contactDto))
	if fillErr != nil {
		panic(fillErr)
	}
	service.contactRepo.Save(&contact)
	return contact, nil
}

func (service *contactService) DeleteById(contactId uint) *gorm.DB {
	return service.contactRepo.DeleteById(contactId)
}
