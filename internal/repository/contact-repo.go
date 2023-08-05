package repository

import (
	"errors"
	"github.com/arief-hidayat/gin-gorm-api/internal/model"

	"gorm.io/gorm"
)

type ContactRepo interface {
	AllContacts() []model.Contact
	Insert(contact model.Contact) model.Contact
	GetById(id uint) (model.Contact, error)
	Save(contact *model.Contact)
	DeleteById(id uint) *gorm.DB
}

type contactRepo struct {
	writerDb *gorm.DB
	readerDb *gorm.DB
}

func NewContactRepo(writerDb *gorm.DB, readerDb *gorm.DB) *contactRepo {
	return &contactRepo{writerDb: writerDb, readerDb: readerDb}
}

func (repo *contactRepo) AllContacts() []model.Contact {
	contacts := []model.Contact{}
	repo.readerDb.Order("id desc").Find(&contacts)
	return contacts
}

func (repo *contactRepo) Insert(contact model.Contact) model.Contact {
	repo.writerDb.Create(&contact)
	return contact
}

func (repo *contactRepo) GetById(id uint) (model.Contact, error) {
	contact := model.Contact{}
	err := repo.readerDb.First(&contact, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return contact, err
	}
	return contact, nil
}

func (repo *contactRepo) Save(contact *model.Contact) {
	repo.writerDb.Save(contact)
}

func (repo *contactRepo) DeleteById(id uint) *gorm.DB {
	return repo.writerDb.Delete(&model.Contact{}, id)
}
