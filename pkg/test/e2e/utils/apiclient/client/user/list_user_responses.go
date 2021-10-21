// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListUserReader is a Reader for the ListUser structure.
type ListUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListUserOK creates a ListUserOK with default headers values
func NewListUserOK() *ListUserOK {
	return &ListUserOK{}
}

/* ListUserOK describes a response with status code 200, with default header values.

User
*/
type ListUserOK struct {
	Payload []*models.User
}

func (o *ListUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserOK  %+v", 200, o.Payload)
}
func (o *ListUserOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *ListUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserUnauthorized creates a ListUserUnauthorized with default headers values
func NewListUserUnauthorized() *ListUserUnauthorized {
	return &ListUserUnauthorized{}
}

/* ListUserUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListUserUnauthorized struct {
}

func (o *ListUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserUnauthorized ", 401)
}

func (o *ListUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUserForbidden creates a ListUserForbidden with default headers values
func NewListUserForbidden() *ListUserForbidden {
	return &ListUserForbidden{}
}

/* ListUserForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListUserForbidden struct {
}

func (o *ListUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUserForbidden ", 403)
}

func (o *ListUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUserDefault creates a ListUserDefault with default headers values
func NewListUserDefault(code int) *ListUserDefault {
	return &ListUserDefault{
		_statusCode: code,
	}
}

/* ListUserDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListUserDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list user default response
func (o *ListUserDefault) Code() int {
	return o._statusCode
}

func (o *ListUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/users][%d] listUser default  %+v", o._statusCode, o.Payload)
}
func (o *ListUserDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}