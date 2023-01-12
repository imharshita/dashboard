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

// ListKubeVirtStorageClassesReader is a Reader for the ListKubeVirtStorageClasses structure.
type ListKubeVirtStorageClassesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubeVirtStorageClassesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubeVirtStorageClassesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubeVirtStorageClassesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubeVirtStorageClassesOK creates a ListKubeVirtStorageClassesOK with default headers values
func NewListKubeVirtStorageClassesOK() *ListKubeVirtStorageClassesOK {
	return &ListKubeVirtStorageClassesOK{}
}

/*
ListKubeVirtStorageClassesOK describes a response with status code 200, with default header values.

StorageClassList
*/
type ListKubeVirtStorageClassesOK struct {
	Payload models.StorageClassList
}

// IsSuccess returns true when this list kube virt storage classes o k response has a 2xx status code
func (o *ListKubeVirtStorageClassesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list kube virt storage classes o k response has a 3xx status code
func (o *ListKubeVirtStorageClassesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list kube virt storage classes o k response has a 4xx status code
func (o *ListKubeVirtStorageClassesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list kube virt storage classes o k response has a 5xx status code
func (o *ListKubeVirtStorageClassesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list kube virt storage classes o k response a status code equal to that given
func (o *ListKubeVirtStorageClassesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListKubeVirtStorageClassesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubeVirtStorageClassesOK  %+v", 200, o.Payload)
}

func (o *ListKubeVirtStorageClassesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubeVirtStorageClassesOK  %+v", 200, o.Payload)
}

func (o *ListKubeVirtStorageClassesOK) GetPayload() models.StorageClassList {
	return o.Payload
}

func (o *ListKubeVirtStorageClassesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubeVirtStorageClassesDefault creates a ListKubeVirtStorageClassesDefault with default headers values
func NewListKubeVirtStorageClassesDefault(code int) *ListKubeVirtStorageClassesDefault {
	return &ListKubeVirtStorageClassesDefault{
		_statusCode: code,
	}
}

/*
ListKubeVirtStorageClassesDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListKubeVirtStorageClassesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list kube virt storage classes default response
func (o *ListKubeVirtStorageClassesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list kube virt storage classes default response has a 2xx status code
func (o *ListKubeVirtStorageClassesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list kube virt storage classes default response has a 3xx status code
func (o *ListKubeVirtStorageClassesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list kube virt storage classes default response has a 4xx status code
func (o *ListKubeVirtStorageClassesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list kube virt storage classes default response has a 5xx status code
func (o *ListKubeVirtStorageClassesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list kube virt storage classes default response a status code equal to that given
func (o *ListKubeVirtStorageClassesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListKubeVirtStorageClassesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubeVirtStorageClasses default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubeVirtStorageClassesDefault) String() string {
	return fmt.Sprintf("[GET /api/v2/providers/kubevirt/storageclasses][%d] listKubeVirtStorageClasses default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubeVirtStorageClassesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListKubeVirtStorageClassesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
