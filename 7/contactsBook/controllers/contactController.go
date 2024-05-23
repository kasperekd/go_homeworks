package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request. Please provide valid JSON data."))
		return
	}
	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	if data == nil {
		u.Respond(w, u.Message(false, "No contacts found"))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	vars := mux.Vars(r)
	contactId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid contact ID"))
		return
	}

	contact := models.GetContact(uint(contactId))
	if contact == nil {
		u.Respond(w, u.Message(false, "Contact not found"))
		return
	}

	if contact.UserId != user {
		u.Respond(w, u.Message(false, "You are not authorized to update this contact"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request. Please provide valid JSON data."))
		return
	}

	resp := contact.UpdateContact()
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	vars := mux.Vars(r)
	contactId, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid contact ID"))
		return
	}

	contact := models.GetContact(uint(contactId))
	if contact == nil {
		u.Respond(w, u.Message(false, "Contact not found"))
		return
	}

	if contact.UserId != user {
		u.Respond(w, u.Message(false, "You are not authorized to delete this contact"))
		return
	}

	resp := contact.DeleteContact()
	u.Respond(w, resp)
}
