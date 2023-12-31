// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// User user
// Example: {"id":"5303d74c64f66366f00cb9b2a94f3251bf5","imageUrl":"https://images.medium.com/0*fkfQiTzT7TlUGGyI.png","name":"Jamie Talbot","url":"https://medium.com/@majelbstoat","username":"majelbstoat"}
//
// swagger:model User
type User struct {

	// A unique identifier for the user.
	ID string `json:"id,omitempty"`

	// The URL to the user’s avatar on Medium
	ImageURL string `json:"imageUrl,omitempty"`

	// The user’s name on Medium.
	Name string `json:"name,omitempty"`

	// The URL to the user’s profile on Medium
	URL string `json:"url,omitempty"`

	// The user’s username on Medium.
	Username string `json:"username,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user based on context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
