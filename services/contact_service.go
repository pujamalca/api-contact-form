// Package services provides business logic implementations for contact-related operations
// in the API Contact Form application.
//
// It defines the ContactService interface and its implementation, which handle the creation,
// retrieval, updating, and deletion of contact records by interacting with the ContactRepository.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package services

import (
	"api-contact-form/models"
	"api-contact-form/repositories"
	"api-contact-form/requests"

	"github.com/go-playground/validator/v10"
)

// ContactService defines the business logic interface for contact operations.
type ContactService interface {
	// CreateContact creates a new contact based on the provided request.
	CreateContact(req *requests.ContactRequest) (*models.Contact, error)
	// GetAllContacts retrieves all non-deleted contacts.
	GetAllContacts() ([]models.Contact, error)
	// GetContactByID retrieves a single contact by its ID.
	GetContactByID(id uint) (*models.Contact, error)
	// UpdateContact updates an existing contact identified by its ID.
	UpdateContact(id uint, req *requests.ContactRequest) (*models.Contact, error)
	// DeleteContact marks a contact as deleted based on its ID.
	DeleteContact(id uint) error
}

// contactService is the concrete implementation of ContactService.
// It interacts with the ContactRepository to perform data operations and uses
// a validator to ensure request data integrity.
type contactService struct {
	repository repositories.ContactRepository
	validate   *validator.Validate
}

// NewContactService creates a new instance of ContactService with the provided ContactRepository.
// It initializes the validator for request validation.
func NewContactService(repository repositories.ContactRepository) ContactService {
	return &contactService{
		repository: repository,
		validate:   validator.New(),
	}
}

// CreateContact creates a new contact based on the provided ContactRequest.
// It validates the request, maps it to the Contact model, and persists it using the repository.
// Returns the created Contact and any error encountered.
func (s *contactService) CreateContact(req *requests.ContactRequest) (*models.Contact, error) {
	// Validate input
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// Map request to Contact model
	contact := models.Contact{
		FullName: req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Message:  req.Message,
	}

	// Persist the contact using the repository
	err := s.repository.Create(&contact)
	return &contact, err
}

// GetAllContacts retrieves all non-deleted contacts from the repository.
// Returns a slice of Contact models and any error encountered.
func (s *contactService) GetAllContacts() ([]models.Contact, error) {
	return s.repository.FindAll()
}

// GetContactByID retrieves a single contact by its ID.
// Returns the Contact model and any error encountered if the contact is not found.
func (s *contactService) GetContactByID(id uint) (*models.Contact, error) {
	return s.repository.FindByID(id)
}

// UpdateContact updates an existing contact identified by its ID based on the provided ContactRequest.
// It validates the request, retrieves the existing contact, updates its fields, and persists the changes.
// Returns the updated Contact and any error encountered.
func (s *contactService) UpdateContact(id uint, req *requests.ContactRequest) (*models.Contact, error) {
	// Validate input
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// Retrieve the existing contact
	contact, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update contact fields
	contact.FullName = req.Name
	contact.Email = req.Email
	contact.Phone = req.Phone
	contact.Message = req.Message

	// Persist the updated contact using the repository
	err = s.repository.Update(contact)
	return contact, err
}

// DeleteContact marks a contact as deleted based on its ID.
// It retrieves the contact and sets its DeletedAt field to the current time.
// Returns any error encountered during the operation.
func (s *contactService) DeleteContact(id uint) error {
	// Retrieve the contact to be deleted
	contact, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	// Mark the contact as deleted
	return s.repository.Delete(contact)
}