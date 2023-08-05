package serializer

import "github.com/arief-hidayat/gin-gorm-api/internal/model"

type ContactResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	MobileNo    string `json:"mobile_no"`
	Institution string `json:"institution"`
}

type ContactSerializer struct {
	Contact model.Contact
}

func (serializer *ContactSerializer) Response() ContactResponse {
	return ContactResponse{
		ID:          serializer.Contact.ID,
		Name:        serializer.Contact.Name,
		Email:       serializer.Contact.Email,
		MobileNo:    serializer.Contact.MobileNo,
		Institution: serializer.Contact.Institution,
	}
}

type ContactsSerializer struct {
	Contacts []model.Contact
}

func (serializer *ContactsSerializer) Response() []ContactResponse {
	var response []ContactResponse
	for _, contact := range serializer.Contacts {
		contactSerializer := ContactSerializer{Contact: contact}
		response = append(response, contactSerializer.Response())
	}
	return response
}
