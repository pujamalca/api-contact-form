// Package models defines the data models for the API Contact Form application.
//
// It includes the Contact struct, which represents a contact message submitted
// through the API. The struct is configured for use with GORM, an ORM library
// for Go, to handle database interactions.
//
// Author: Tri Wicaksono
// Website: https://triwicaksono.com
package models

import (
	"time"
)

// Contact represents a contact message submitted through the API.
type Contact struct {
	// ID is the unique identifier for each contact message.
	ID uint `gorm:"primaryKey;column:id;type:BIGINT UNSIGNED AUTO_INCREMENT"`

	// FullName is the name of the person submitting the contact message.
	FullName string `gorm:"column:full_name;type:VARCHAR(100);not null"`

	// Email is the email address of the person submitting the contact message.
	Email string `gorm:"column:email_address;type:VARCHAR(100);not null"`

	// Phone is the phone number of the person submitting the contact message.
	Phone string `gorm:"column:phone_number;type:VARCHAR(20);not null"`

	// Message is the content of the contact message.
	Message string `gorm:"column:message_text;type:TEXT;not null"`

	// CreatedAt records the timestamp when the contact message was created.
	CreatedAt time.Time `gorm:"column:created_at;type:DATETIME;autoCreateTime"`

	// UpdatedAt records the timestamp when the contact message was last updated.
	UpdatedAt time.Time `gorm:"column:updated_at;type:DATETIME;autoUpdateTime"`

	// DeletedAt records the timestamp when the contact message was deleted.
	// This field is indexed to optimize deletion queries.
	DeletedAt time.Time `gorm:"column:deleted_at;type:DATETIME;index"`
}

// TableName specifies the table name for the Contact model in the database.
func (Contact) TableName() string {
	return "contact_messages"
}