// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/dashboard/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListProjectKubeVirtPreferencesReader is a Reader for the ListProjectKubeVirtPreferences structure.
type ListProjectKubeVirtPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectKubeVirtPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectKubeVirtPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProjectKubeVirtPreferencesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProjectKubeVirtPreferencesOK creates a ListProjectKubeVirtPreferencesOK with default headers values
func NewListProjectKubeVirtPreferencesOK() *ListProjectKubeVirtPreferencesOK {
	return &ListProjectKubeVirtPreferencesOK{}
}

/*
ListProjectKubeVirtPreferencesOK describes a response with status code 200, with default header values.

VirtualMachinePreferenceList
*/
type ListProjectKubeVirtPreferencesOK struct {
	Payload *models.VirtualMachinePreferenceList
}

// IsSuccess returns true when this list project kube virt preferences o k response has a 2xx status code
func (o *ListProjectKubeVirtPreferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project kube virt preferences o k response has a 3xx status code
func (o *ListProjectKubeVirtPreferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project kube virt preferences o k response has a 4xx status code
func (o *ListProjectKubeVirtPreferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project kube virt preferences o k response has a 5xx status code
func (o *ListProjectKubeVirtPreferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project kube virt preferences o k response a status code equal to that given
func (o *ListProjectKubeVirtPreferencesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListProjectKubeVirtPreferencesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/preferences][%d] listProjectKubeVirtPreferencesOK  %+v", 200, o.Payload)
}

func (o *ListProjectKubeVirtPreferencesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/preferences][%d] listProjectKubeVirtPreferencesOK  %+v", 200, o.Payload)
}

func (o *ListProjectKubeVirtPreferencesOK) GetPayload() *models.VirtualMachinePreferenceList {
	return o.Payload
}

func (o *ListProjectKubeVirtPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachinePreferenceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectKubeVirtPreferencesDefault creates a ListProjectKubeVirtPreferencesDefault with default headers values
func NewListProjectKubeVirtPreferencesDefault(code int) *ListProjectKubeVirtPreferencesDefault {
	return &ListProjectKubeVirtPreferencesDefault{
		_statusCode: code,
	}
}

/*
ListProjectKubeVirtPreferencesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListProjectKubeVirtPreferencesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list project kube virt preferences default response
func (o *ListProjectKubeVirtPreferencesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list project kube virt preferences default response has a 2xx status code
func (o *ListProjectKubeVirtPreferencesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list project kube virt preferences default response has a 3xx status code
func (o *ListProjectKubeVirtPreferencesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list project kube virt preferences default response has a 4xx status code
func (o *ListProjectKubeVirtPreferencesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list project kube virt preferences default response has a 5xx status code
func (o *ListProjectKubeVirtPreferencesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list project kube virt preferences default response a status code equal to that given
func (o *ListProjectKubeVirtPreferencesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListProjectKubeVirtPreferencesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/preferences][%d] listProjectKubeVirtPreferences default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectKubeVirtPreferencesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/providers/kubevirt/preferences][%d] listProjectKubeVirtPreferences default  %+v", o._statusCode, o.Payload)
}

func (o *ListProjectKubeVirtPreferencesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListProjectKubeVirtPreferencesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
