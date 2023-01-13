// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectOpenstackTenantsReader is a Reader for the ListProjectOpenstackTenants structure.
type ListProjectOpenstackTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectOpenstackTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectOpenstackTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectOpenstackTenantsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectOpenstackTenantsOK creates a ListProjectOpenstackTenantsOK with default headers values
func NewListProjectOpenstackTenantsOK() *ListProjectOpenstackTenantsOK {
	return &ListProjectOpenstackTenantsOK{}
}

/*
ListProjectOpenstackTenantsOK describes a response with status code 200, with default header values.

OpenstackTenant
*/
type ListProjectOpenstackTenantsOK struct {
	Payload []*models.OpenstackTenant
}

// IsSuccess returns true when this list project openstack tenants o k response has a 2xx status code
func (o *ListProjectOpenstackTenantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project openstack tenants o k response has a 3xx status code
func (o *ListProjectOpenstackTenantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project openstack tenants o k response has a 4xx status code
func (o *ListProjectOpenstackTenantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project openstack tenants o k response has a 5xx status code
func (o *ListProjectOpenstackTenantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project openstack tenants o k response a status code equal to that given
func (o *ListProjectOpenstackTenantsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectOpenstackTenantsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/openstack/tenants][%d] listProjectOpenstackTenantsOK  %+v", 200, o.Payload)
}

func (o *ListProjectOpenstackTenantsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/openstack/tenants][%d] listProjectOpenstackTenantsOK  %+v", 200, o.Payload)
}

func (o *ListProjectOpenstackTenantsOK) GetPayload() []*models.OpenstackTenant {
	return o.Payload
}

func (o *ListProjectOpenstackTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectOpenstackTenantsDefault creates a ListProjectOpenstackTenantsDefault with default headers values
func NewListProjectOpenstackTenantsDefault(code int) *ListProjectOpenstackTenantsDefault {
	return &ListProjectOpenstackTenantsDefault{
		_statusCode: code,
	}
}

/*
ListProjectOpenstackTenantsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectOpenstackTenantsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project openstack tenants default response
func (o *ListProjectOpenstackTenantsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project openstack tenants default response has a 2xx status code
func (o *ListProjectOpenstackTenantsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project openstack tenants default response has a 3xx status code
func (o *ListProjectOpenstackTenantsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project openstack tenants default response has a 4xx status code
func (o *ListProjectOpenstackTenantsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project openstack tenants default response has a 5xx status code
func (o *ListProjectOpenstackTenantsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project openstack tenants default response a status code equal to that given
func (o *ListProjectOpenstackTenantsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectOpenstackTenantsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/openstack/tenants][%d] listProjectOpenstackTenants default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectOpenstackTenantsDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/openstack/tenants][%d] listProjectOpenstackTenants default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectOpenstackTenantsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectOpenstackTenantsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
