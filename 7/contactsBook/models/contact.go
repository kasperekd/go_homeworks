package models

import (
	u "contactsBook/utils"
	"fmt"
	"regexp"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Name cannot be empty"), false
	}
	if contact.Phone == "" {
		return u.Message(false, "Phone number cannot be empty"), false
	}
	if !isValidPhone(contact.Phone) {
		return u.Message(false, "Phone number is not valid"), false
	}
	if contact.UserId <= 0 {
		return u.Message(false, "User not found"), false
	}
	return u.Message(true, "success"), true
}

func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return re.MatchString(phone)
}

func (contact *Contact) CreateContact() map[string]interface{} {
	if response, ok := contact.ValidateContact(); !ok {
		return response
	}
	GetDB().Create(contact)
	resp := u.Message(true, "Contact created successfully")
	resp["contact"] = contact
	return resp
}

func (contact *Contact) UpdateContact() map[string]interface{} {
	if response, ok := contact.ValidateContact(); !ok {
		return response
	}
	GetDB().Save(contact)
	resp := u.Message(true, "Contact updated successfully")
	resp["contact"] = contact
	return resp
}

func (contact *Contact) DeleteContact() map[string]interface{} {
	GetDB().Delete(contact)
	resp := u.Message(true, "Contact deleted successfully")
	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().First(contact, id).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}
