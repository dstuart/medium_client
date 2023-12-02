// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContibutorResponse Contributors list for a publication
//
// list of contributors for a given publication
// Example: {"data":[{"publicationId":"b45573563f5a","role":"editor","userId":"13a06af8f81849c64dafbce822cbafbfab7ed7cecf82135bca946807ea351290d"},{"publicationId":"b45573563f5a","role":"editor","userId":"1c9c63b15b874d3e354340b7d7458d55e1dda0f6470074df1cc99608a372866ac"},{"publicationId":"b45573563f5a","role":"editor","userId":"1cc07499453463518b77d31650c0b53609dc973ad8ebd33690c7be9236e9384ad"},{"publicationId":"b45573563f5a","role":"writer","userId":"196f70942410555f4b3030debc4f199a0d5a0309a7b9df96c57b8ec6e4b5f11d7"},{"publicationId":"b45573563f5a","role":"writer","userId":"14d4a581f21ff537d245461b8ff2ae9b271b57d9554e25d863e3df6ef03ddd480"}]}
//
// swagger:model ContibutorResponse
type ContibutorResponse struct {

	// data
	Data []*Contibutor `json:"data"`
}

// Validate validates this contibutor response
func (m *ContibutorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContibutorResponse) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this contibutor response based on the context it is used
func (m *ContibutorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContibutorResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {

			if swag.IsZero(m.Data[i]) { // not required
				return nil
			}

			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContibutorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContibutorResponse) UnmarshalBinary(b []byte) error {
	var res ContibutorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
