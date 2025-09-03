// Package requests defines the request payload structures for the API Contact Form application.
//
// It includes the ContactRequest struct, which represents the data required to create or update
// a contact message through the API.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package requests

// ContactRequest represents the payload for creating or updating a contact message.
type ContactRequest struct {
	// Name is the full name of the person submitting the contact message.
	// It is a required field with a maximum length of 100 characters.
	Name string `json:"name" binding:"required,max=100"`

	// Email is the email address of the person submitting the contact message.
	// It is a required field with a maximum length of 100 characters and must follow a valid email format.
	Email string `json:"email" binding:"required,email,max=100"`

	// Phone is the phone number of the person submitting the contact message.
	// It is a required field with a maximum length of 20 characters.
	Phone string `json:"phone" binding:"required,max=20"`

	// Message is the content of the contact message.
	// It is a required field.
	Message string `json:"message" binding:"required"`
}