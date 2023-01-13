// Code generated by go-swagger; DO NOT EDIT.

package digitalocean

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectDigitaloceanSizesReader is a Reader for the ListProjectDigitaloceanSizes structure.
type ListProjectDigitaloceanSizesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectDigitaloceanSizesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectDigitaloceanSizesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectDigitaloceanSizesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectDigitaloceanSizesOK creates a ListProjectDigitaloceanSizesOK with default headers values
func NewListProjectDigitaloceanSizesOK() *ListProjectDigitaloceanSizesOK {
	return &ListProjectDigitaloceanSizesOK{}
}

/*
ListProjectDigitaloceanSizesOK describes a response with status code 200, with default header values.

DigitaloceanSizeList
*/
type ListProjectDigitaloceanSizesOK struct {
	Payload *models.DigitaloceanSizeList
}

// IsSuccess returns true when this list project digitalocean sizes o k response has a 2xx status code
func (o *ListProjectDigitaloceanSizesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project digitalocean sizes o k response has a 3xx status code
func (o *ListProjectDigitaloceanSizesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project digitalocean sizes o k response has a 4xx status code
func (o *ListProjectDigitaloceanSizesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project digitalocean sizes o k response has a 5xx status code
func (o *ListProjectDigitaloceanSizesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project digitalocean sizes o k response a status code equal to that given
func (o *ListProjectDigitaloceanSizesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectDigitaloceanSizesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/digitalocean/sizes][%d] listProjectDigitaloceanSizesOK  %+v", 200, o.Payload)
}

func (o *ListProjectDigitaloceanSizesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/digitalocean/sizes][%d] listProjectDigitaloceanSizesOK  %+v", 200, o.Payload)
}

func (o *ListProjectDigitaloceanSizesOK) GetPayload() *models.DigitaloceanSizeList {
	return o.Payload
}

func (o *ListProjectDigitaloceanSizesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DigitaloceanSizeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectDigitaloceanSizesDefault creates a ListProjectDigitaloceanSizesDefault with default headers values
func NewListProjectDigitaloceanSizesDefault(code int) *ListProjectDigitaloceanSizesDefault {
	return &ListProjectDigitaloceanSizesDefault{
		_statusCode: code,
	}
}

/*
ListProjectDigitaloceanSizesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectDigitaloceanSizesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project digitalocean sizes default response
func (o *ListProjectDigitaloceanSizesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project digitalocean sizes default response has a 2xx status code
func (o *ListProjectDigitaloceanSizesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project digitalocean sizes default response has a 3xx status code
func (o *ListProjectDigitaloceanSizesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project digitalocean sizes default response has a 4xx status code
func (o *ListProjectDigitaloceanSizesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project digitalocean sizes default response has a 5xx status code
func (o *ListProjectDigitaloceanSizesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project digitalocean sizes default response a status code equal to that given
func (o *ListProjectDigitaloceanSizesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectDigitaloceanSizesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/digitalocean/sizes][%d] listProjectDigitaloceanSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectDigitaloceanSizesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/digitalocean/sizes][%d] listProjectDigitaloceanSizes default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectDigitaloceanSizesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectDigitaloceanSizesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
