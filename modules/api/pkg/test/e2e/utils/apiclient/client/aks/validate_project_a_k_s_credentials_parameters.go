// Code generated by go-swagger; DO NOT EDIT.

package aks

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

// NewValidateProjectAKSCredentialsParams creates a new ValidateProjectAKSCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateProjectAKSCredentialsParams() *ValidateProjectAKSCredentialsParams {
	return &ValidateProjectAKSCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateProjectAKSCredentialsParamsWithTimeout creates a new ValidateProjectAKSCredentialsParams object
// with the ability to set a timeout on a request.
func NewValidateProjectAKSCredentialsParamsWithTimeout(timeout time.Duration) *ValidateProjectAKSCredentialsParams {
	return &ValidateProjectAKSCredentialsParams{
		timeout: timeout,
	}
}

// NewValidateProjectAKSCredentialsParamsWithContext creates a new ValidateProjectAKSCredentialsParams object
// with the ability to set a context for a request.
func NewValidateProjectAKSCredentialsParamsWithContext(ctx context.Context) *ValidateProjectAKSCredentialsParams {
	return &ValidateProjectAKSCredentialsParams{
		Context: ctx,
	}
}

// NewValidateProjectAKSCredentialsParamsWithHTTPClient creates a new ValidateProjectAKSCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateProjectAKSCredentialsParamsWithHTTPClient(client *http.Client) *ValidateProjectAKSCredentialsParams {
	return &ValidateProjectAKSCredentialsParams{
		HTTPClient: client,
	}
}

/*
ValidateProjectAKSCredentialsParams contains all the parameters to send to the API endpoint

	for the validate project a k s credentials operation.

	Typically these are written to a http.Request.
*/
type ValidateProjectAKSCredentialsParams struct {

	// ClientID.
	ClientID *string

	// ClientSecret.
	ClientSecret *string

	// Credential.
	Credential *string

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

// WithDefaults hydrates default values in the validate project a k s credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateProjectAKSCredentialsParams) WithDefaults() *ValidateProjectAKSCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate project a k s credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateProjectAKSCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithTimeout(timeout time.Duration) *ValidateProjectAKSCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithContext(ctx context.Context) *ValidateProjectAKSCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithHTTPClient(client *http.Client) *ValidateProjectAKSCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithClientID(clientID *string) *ValidateProjectAKSCredentialsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithClientSecret(clientSecret *string) *ValidateProjectAKSCredentialsParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithCredential adds the credential to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithCredential(credential *string) *ValidateProjectAKSCredentialsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithSubscriptionID adds the subscriptionID to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithSubscriptionID(subscriptionID *string) *ValidateProjectAKSCredentialsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetSubscriptionID(subscriptionID *string) {
	o.SubscriptionID = subscriptionID
}

// WithTenantID adds the tenantID to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithTenantID(tenantID *string) *ValidateProjectAKSCredentialsParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithProjectID adds the projectID to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) WithProjectID(projectID string) *ValidateProjectAKSCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the validate project a k s credentials params
func (o *ValidateProjectAKSCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateProjectAKSCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
