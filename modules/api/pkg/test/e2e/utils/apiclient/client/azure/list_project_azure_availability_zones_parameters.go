// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewListProjectAzureAvailabilityZonesParams creates a new ListProjectAzureAvailabilityZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectAzureAvailabilityZonesParams() *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectAzureAvailabilityZonesParamsWithTimeout creates a new ListProjectAzureAvailabilityZonesParams object
// with the ability to set a timeout on a request.
func NewListProjectAzureAvailabilityZonesParamsWithTimeout(timeout time.Duration) *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		timeout: timeout,
	}
}

// NewListProjectAzureAvailabilityZonesParamsWithContext creates a new ListProjectAzureAvailabilityZonesParams object
// with the ability to set a context for a request.
func NewListProjectAzureAvailabilityZonesParamsWithContext(ctx context.Context) *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		Context: ctx,
	}
}

// NewListProjectAzureAvailabilityZonesParamsWithHTTPClient creates a new ListProjectAzureAvailabilityZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectAzureAvailabilityZonesParamsWithHTTPClient(client *http.Client) *ListProjectAzureAvailabilityZonesParams {
	return &ListProjectAzureAvailabilityZonesParams{
		HTTPClient: client,
	}
}

/*
ListProjectAzureAvailabilityZonesParams contains all the parameters to send to the API endpoint

	for the list project azure availability zones operation.

	Typically these are written to a http.Request.
*/
type ListProjectAzureAvailabilityZonesParams struct {

	// ClientID.
	ClientID *string

	// ClientSecret.
	ClientSecret *string

	// Credential.
	Credential *string

	// Location.
	Location *string

	// SKUName.
	SKUName *string

	// SubscriptionID.
	SubscriptionID *string

	// TenantID.
	TenantID *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project azure availability zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAzureAvailabilityZonesParams) WithDefaults() *ListProjectAzureAvailabilityZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project azure availability zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectAzureAvailabilityZonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithTimeout(timeout time.Duration) *ListProjectAzureAvailabilityZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithContext(ctx context.Context) *ListProjectAzureAvailabilityZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithHTTPClient(client *http.Client) *ListProjectAzureAvailabilityZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithClientID(clientID *string) *ListProjectAzureAvailabilityZonesParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithClientSecret(clientSecret *string) *ListProjectAzureAvailabilityZonesParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithCredential adds the credential to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithCredential(credential *string) *ListProjectAzureAvailabilityZonesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithLocation adds the location to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithLocation(location *string) *ListProjectAzureAvailabilityZonesParams {
	o.SetLocation(location)
	return o
}

// SetLocation adds the location to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetLocation(location *string) {
	o.Location = location
}

// WithSKUName adds the sKUName to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithSKUName(sKUName *string) *ListProjectAzureAvailabilityZonesParams {
	o.SetSKUName(sKUName)
	return o
}

// SetSKUName adds the sKUName to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetSKUName(sKUName *string) {
	o.SKUName = sKUName
}

// WithSubscriptionID adds the subscriptionID to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithSubscriptionID(subscriptionID *string) *ListProjectAzureAvailabilityZonesParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WithTenantID adds the tenantID to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithTenantID(tenantID *string) *ListProjectAzureAvailabilityZonesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithProjectID adds the projectID to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) WithProjectID(projectID string) *ListProjectAzureAvailabilityZonesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project azure availability zones params
func (o *ListProjectAzureAvailabilityZonesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectAzureAvailabilityZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientID != nil {

		// header param ClientID
		if err := r.SetHeaderParam("ClientID", *o.ClientID); err != nil {
			return err
		}
	}

	if o.ClientSecret != nil {

		// header param ClientSecret
		if err := r.SetHeaderParam("ClientSecret", *o.ClientSecret); err != nil {
			return err
		}
	}

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Location != nil {

		// header param Location
		if err := r.SetHeaderParam("Location", *o.Location); err != nil {
			return err
		}
	}

	if o.SKUName != nil {

		// header param SKUName
		if err := r.SetHeaderParam("SKUName", *o.SKUName); err != nil {
			return err
		}
	}

	if o.SubscriptionID != nil {

		// header param SubscriptionID
		if err := r.SetHeaderParam("SubscriptionID", *o.SubscriptionID); err != nil {
			return err
		}
	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
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
