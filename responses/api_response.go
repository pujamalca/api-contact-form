// Package responses defines the response payload structures for the API Contact Form application.
//
// It includes the APIResponse struct for standard API responses and the ContactResponse struct
// for representing contact data in responses. Additionally, it provides helper functions
// to convert models to response formats.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package responses

import (
	"api-contact-form/helpers"
	"api-contact-form/models"
)

// APIResponse represents the standard structure for API responses.
type APIResponse struct {
	// Code is a string representing the status code of the response.
	Code string `json:"code"`
	// Message provides a human-readable message about the response.
	Message string `json:"message"`
	// Data holds the payload of the response, which can be any type.
	Data interface{} `json:"data"`
}

// ContactResponse represents the structure of a contact in API responses.
type ContactResponse struct {
	// ID is the unique identifier of the contact.
	ID uint `json:"id"`
	// Name is the full name of the contact.
	Name string `json:"name"`
	// Email is the email address of the contact.
	Email string `json:"email"`
	// Phone is the phone number of the contact.
	Phone string `json:"phone"`
	// Message is the message content provided by the contact.
	Message string `json:"message"`
	// CreatedAt is the timestamp when the contact was created, formatted as a human-readable string.
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the timestamp when the contact was last updated, formatted as a human-readable string.
	UpdatedAt string `json:"updated_at"`
}

// ContactResponseFromModel converts a Contact model to a ContactResponse.
//
// Parameters:
//   - contact: A pointer to the Contact model to be converted.
//
// Returns:
//   - A ContactResponse struct populated with data from the Contact model.
func ContactResponseFromModel(contact *models.Contact) ContactResponse {
	return ContactResponse{
		ID:        contact.ID,
		Name:      contact.FullName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		Message:   contact.Message,
		CreatedAt: helpers.FormatTimeHuman(contact.CreatedAt),
		UpdatedAt: helpers.FormatTimeHuman(contact.UpdatedAt),
	}
}