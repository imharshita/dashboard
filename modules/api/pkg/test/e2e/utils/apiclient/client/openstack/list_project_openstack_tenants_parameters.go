// Code generated by go-swagger; DO NOT EDIT.

package openstack

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
	"github.com/go-openapi/swag"
)

// NewListProjectOpenstackTenantsParams creates a new ListProjectOpenstackTenantsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectOpenstackTenantsParams() *ListProjectOpenstackTenantsParams {
	return &ListProjectOpenstackTenantsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectOpenstackTenantsParamsWithTimeout creates a new ListProjectOpenstackTenantsParams object
// with the ability to set a timeout on a request.
func NewListProjectOpenstackTenantsParamsWithTimeout(timeout time.Duration) *ListProjectOpenstackTenantsParams {
	return &ListProjectOpenstackTenantsParams{
		timeout: timeout,
	}
}

// NewListProjectOpenstackTenantsParamsWithContext creates a new ListProjectOpenstackTenantsParams object
// with the ability to set a context for a request.
func NewListProjectOpenstackTenantsParamsWithContext(ctx context.Context) *ListProjectOpenstackTenantsParams {
	return &ListProjectOpenstackTenantsParams{
		Context: ctx,
	}
}

// NewListProjectOpenstackTenantsParamsWithHTTPClient creates a new ListProjectOpenstackTenantsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectOpenstackTenantsParamsWithHTTPClient(client *http.Client) *ListProjectOpenstackTenantsParams {
	return &ListProjectOpenstackTenantsParams{
		HTTPClient: client,
	}
}

/*
ListProjectOpenstackTenantsParams contains all the parameters to send to the API endpoint

	for the list project openstack tenants operation.

	Typically these are written to a http.Request.
*/
type ListProjectOpenstackTenantsParams struct {

	// ApplicationCredentialID.
	ApplicationCredentialID *string

	// ApplicationCredentialSecret.
	ApplicationCredentialSecret *string

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// Domain.
	Domain *string

	// OIDCAuthentication.
	OIDCAuthentication *bool

	// Password.
	Password *string

	// Username.
	Username *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project openstack tenants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectOpenstackTenantsParams) WithDefaults() *ListProjectOpenstackTenantsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project openstack tenants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectOpenstackTenantsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithTimeout(timeout time.Duration) *ListProjectOpenstackTenantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithContext(ctx context.Context) *ListProjectOpenstackTenantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithHTTPClient(client *http.Client) *ListProjectOpenstackTenantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationCredentialID adds the applicationCredentialID to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithApplicationCredentialID(applicationCredentialID *string) *ListProjectOpenstackTenantsParams {
	o.SetApplicationCredentialID(applicationCredentialID)
	return o
}

// SetApplicationCredentialID adds the applicationCredentialId to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetApplicationCredentialID(applicationCredentialID *string) {
	o.ApplicationCredentialID = applicationCredentialID
}

// WithApplicationCredentialSecret adds the applicationCredentialSecret to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithApplicationCredentialSecret(applicationCredentialSecret *string) *ListProjectOpenstackTenantsParams {
	o.SetApplicationCredentialSecret(applicationCredentialSecret)
	return o
}

// SetApplicationCredentialSecret adds the applicationCredentialSecret to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetApplicationCredentialSecret(applicationCredentialSecret *string) {
	o.ApplicationCredentialSecret = applicationCredentialSecret
}

// WithCredential adds the credential to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithCredential(credential *string) *ListProjectOpenstackTenantsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithDatacenterName(datacenterName *string) *ListProjectOpenstackTenantsParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithDomain adds the domain to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithDomain(domain *string) *ListProjectOpenstackTenantsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithOIDCAuthentication adds the oIDCAuthentication to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithOIDCAuthentication(oIDCAuthentication *bool) *ListProjectOpenstackTenantsParams {
	o.SetOIDCAuthentication(oIDCAuthentication)
	return o
}

// SetOIDCAuthentication adds the oIdCAuthentication to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetOIDCAuthentication(oIDCAuthentication *bool) {
	o.OIDCAuthentication = oIDCAuthentication
}

// WithPassword adds the password to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithPassword(password *string) *ListProjectOpenstackTenantsParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithUsername(username *string) *ListProjectOpenstackTenantsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetUsername(username *string) {
	o.Username = username
}

// WithProjectID adds the projectID to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) WithProjectID(projectID string) *ListProjectOpenstackTenantsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project openstack tenants params
func (o *ListProjectOpenstackTenantsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectOpenstackTenantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationCredentialID != nil {

		// header param ApplicationCredentialID
		if err := r.SetHeaderParam("ApplicationCredentialID", *o.ApplicationCredentialID); err != nil {
			return err
		}
	}

	if o.ApplicationCredentialSecret != nil {

		// header param ApplicationCredentialSecret
		if err := r.SetHeaderParam("ApplicationCredentialSecret", *o.ApplicationCredentialSecret); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}
	}

	if o.Domain != nil {

		// header param Domain
		if err := r.SetHeaderParam("Domain", *o.Domain); err != nil {
			return err
		}
	}

	if o.OIDCAuthentication != nil {

		// header param OIDCAuthentication
		if err := r.SetHeaderParam("OIDCAuthentication", swag.FormatBool(*o.OIDCAuthentication)); err != nil {
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

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
