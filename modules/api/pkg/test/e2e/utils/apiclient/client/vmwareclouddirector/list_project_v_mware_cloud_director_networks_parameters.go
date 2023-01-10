// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

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

// NewListProjectVMwareCloudDirectorNetworksParams creates a new ListProjectVMwareCloudDirectorNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectVMwareCloudDirectorNetworksParams() *ListProjectVMwareCloudDirectorNetworksParams {
	return &ListProjectVMwareCloudDirectorNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectVMwareCloudDirectorNetworksParamsWithTimeout creates a new ListProjectVMwareCloudDirectorNetworksParams object
// with the ability to set a timeout on a request.
func NewListProjectVMwareCloudDirectorNetworksParamsWithTimeout(timeout time.Duration) *ListProjectVMwareCloudDirectorNetworksParams {
	return &ListProjectVMwareCloudDirectorNetworksParams{
		timeout: timeout,
	}
}

// NewListProjectVMwareCloudDirectorNetworksParamsWithContext creates a new ListProjectVMwareCloudDirectorNetworksParams object
// with the ability to set a context for a request.
func NewListProjectVMwareCloudDirectorNetworksParamsWithContext(ctx context.Context) *ListProjectVMwareCloudDirectorNetworksParams {
	return &ListProjectVMwareCloudDirectorNetworksParams{
		Context: ctx,
	}
}

// NewListProjectVMwareCloudDirectorNetworksParamsWithHTTPClient creates a new ListProjectVMwareCloudDirectorNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectVMwareCloudDirectorNetworksParamsWithHTTPClient(client *http.Client) *ListProjectVMwareCloudDirectorNetworksParams {
	return &ListProjectVMwareCloudDirectorNetworksParams{
		HTTPClient: client,
	}
}

/*
ListProjectVMwareCloudDirectorNetworksParams contains all the parameters to send to the API endpoint

	for the list project v mware cloud director networks operation.

	Typically these are written to a http.Request.
*/
type ListProjectVMwareCloudDirectorNetworksParams struct {

	// Credential.
	Credential *string

	// Organization.
	Organization *string

	// Password.
	Password *string

	// Username.
	Username *string

	// VDC.
	VDC *string

	/* Dc.

	   KKP Datacenter to use for endpoint
	*/
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project v mware cloud director networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithDefaults() *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project v mware cloud director networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithTimeout(timeout time.Duration) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithContext(ctx context.Context) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithHTTPClient(client *http.Client) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithCredential(credential *string) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithOrganization adds the organization to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithOrganization(organization *string) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetOrganization(organization *string) {
	o.Organization = organization
}

// WithPassword adds the password to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithPassword(password *string) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithUsername(username *string) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetUsername(username *string) {
	o.Username = username
}

// WithVDC adds the vDC to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithVDC(vDC *string) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetVDC(vDC)
	return o
}

// SetVDC adds the vDC to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetVDC(vDC *string) {
	o.VDC = vDC
}

// WithDC adds the dc to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithDC(dc string) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) WithProjectID(projectID string) *ListProjectVMwareCloudDirectorNetworksParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project v mware cloud director networks params
func (o *ListProjectVMwareCloudDirectorNetworksParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectVMwareCloudDirectorNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Organization != nil {

		// header param Organization
		if err := r.SetHeaderParam("Organization", *o.Organization); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}
	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
			return err
		}
	}

	if o.VDC != nil {

		// header param VDC
		if err := r.SetHeaderParam("VDC", *o.VDC); err != nil {
			return err
		}
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
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
