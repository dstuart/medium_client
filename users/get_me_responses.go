// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/dstuart/models"
)

// GetMeReader is a Reader for the GetMe structure.
type GetMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /me] GetMe", response, response.Code())
	}
}

// NewGetMeOK creates a GetMeOK with default headers values
func NewGetMeOK() *GetMeOK {
	return &GetMeOK{}
}

/*
GetMeOK describes a response with status code 200, with default header values.

OK
*/
type GetMeOK struct {
	Payload *models.UserResponse
}

// IsSuccess returns true when this get me o k response has a 2xx status code
func (o *GetMeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get me o k response has a 3xx status code
func (o *GetMeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get me o k response has a 4xx status code
func (o *GetMeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get me o k response has a 5xx status code
func (o *GetMeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get me o k response a status code equal to that given
func (o *GetMeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get me o k response
func (o *GetMeOK) Code() int {
	return 200
}

func (o *GetMeOK) Error() string {
	return fmt.Sprintf("[GET /me][%d] getMeOK  %+v", 200, o.Payload)
}

func (o *GetMeOK) String() string {
	return fmt.Sprintf("[GET /me][%d] getMeOK  %+v", 200, o.Payload)
}

func (o *GetMeOK) GetPayload() *models.UserResponse {
	return o.Payload
}

func (o *GetMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeUnauthorized creates a GetMeUnauthorized with default headers values
func NewGetMeUnauthorized() *GetMeUnauthorized {
	return &GetMeUnauthorized{}
}

/*
GetMeUnauthorized describes a response with status code 401, with default header values.

The `accessToken` is invalid or has been revoked.
*/
type GetMeUnauthorized struct {
}

// IsSuccess returns true when this get me unauthorized response has a 2xx status code
func (o *GetMeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get me unauthorized response has a 3xx status code
func (o *GetMeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get me unauthorized response has a 4xx status code
func (o *GetMeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get me unauthorized response has a 5xx status code
func (o *GetMeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get me unauthorized response a status code equal to that given
func (o *GetMeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get me unauthorized response
func (o *GetMeUnauthorized) Code() int {
	return 401
}

func (o *GetMeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /me][%d] getMeUnauthorized ", 401)
}

func (o *GetMeUnauthorized) String() string {
	return fmt.Sprintf("[GET /me][%d] getMeUnauthorized ", 401)
}

func (o *GetMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
