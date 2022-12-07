// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListProjectGCPSubnetworksParams creates a new ListProjectGCPSubnetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectGCPSubnetworksParams() *ListProjectGCPSubnetworksParams {
	return &ListProjectGCPSubnetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectGCPSubnetworksParamsWithTimeout creates a new ListProjectGCPSubnetworksParams object
// with the ability to set a timeout on a request.
func NewListProjectGCPSubnetworksParamsWithTimeout(timeout time.Duration) *ListProjectGCPSubnetworksParams {
	return &ListProjectGCPSubnetworksParams{
		timeout: timeout,
	}
}

// NewListProjectGCPSubnetworksParamsWithContext creates a new ListProjectGCPSubnetworksParams object
// with the ability to set a context for a request.
func NewListProjectGCPSubnetworksParamsWithContext(ctx context.Context) *ListProjectGCPSubnetworksParams {
	return &ListProjectGCPSubnetworksParams{
		Context: ctx,
	}
}

// NewListProjectGCPSubnetworksParamsWithHTTPClient creates a new ListProjectGCPSubnetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectGCPSubnetworksParamsWithHTTPClient(client *http.Client) *ListProjectGCPSubnetworksParams {
	return &ListProjectGCPSubnetworksParams{
		HTTPClient: client,
	}
}

/*
ListProjectGCPSubnetworksParams contains all the parameters to send to the API endpoint

	for the list project g c p subnetworks operation.

	Typically these are written to a http.Request.
*/
type ListProjectGCPSubnetworksParams struct {

	// Credential.
	Credential *string

	// DC.
	DC *string

	// Network.
	Network *string

	// ServiceAccount.
	ServiceAccount *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project g c p subnetworks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectGCPSubnetworksParams) WithDefaults() *ListProjectGCPSubnetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project g c p subnetworks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectGCPSubnetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithTimeout(timeout time.Duration) *ListProjectGCPSubnetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithContext(ctx context.Context) *ListProjectGCPSubnetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithHTTPClient(client *http.Client) *ListProjectGCPSubnetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithCredential(credential *string) *ListProjectGCPSubnetworksParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDC adds the dC to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithDC(dC *string) *ListProjectGCPSubnetworksParams {
	o.SetDC(dC)
	return o
}

// SetDC adds the dC to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetDC(dC *string) {
	o.DC = dC
}

// WithNetwork adds the network to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithNetwork(network *string) *ListProjectGCPSubnetworksParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetNetwork(network *string) {
	o.Network = network
}

// WithServiceAccount adds the serviceAccount to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithServiceAccount(serviceAccount *string) *ListProjectGCPSubnetworksParams {
	o.SetServiceAccount(serviceAccount)
	return o
}

// SetServiceAccount adds the serviceAccount to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetServiceAccount(serviceAccount *string) {
	o.ServiceAccount = serviceAccount
}

// WithProjectID adds the projectID to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) WithProjectID(projectID string) *ListProjectGCPSubnetworksParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project g c p subnetworks params
func (o *ListProjectGCPSubnetworksParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectGCPSubnetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.DC != nil {

		// query param DC
		var qrDC string

		if o.DC != nil {
			qrDC = *o.DC
		}
		qDC := qrDC
		if qDC != "" {

			if err := r.SetQueryParam("DC", qDC); err != nil {
				return err
			}
		}
	}

	if o.Network != nil {

		// query param Network
		var qrNetwork string

		if o.Network != nil {
			qrNetwork = *o.Network
		}
		qNetwork := qrNetwork
		if qNetwork != "" {

			if err := r.SetQueryParam("Network", qNetwork); err != nil {
				return err
			}
		}
	}

	if o.ServiceAccount != nil {

		// header param ServiceAccount
		if err := r.SetHeaderParam("ServiceAccount", *o.ServiceAccount); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
