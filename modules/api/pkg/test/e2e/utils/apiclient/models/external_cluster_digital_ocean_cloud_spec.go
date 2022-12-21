// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExternalClusterDigitalOceanCloudSpec ExternalClusterDigitalOceanCloudSpec specifies access data to DigitalOcean.
//
// swagger:model ExternalClusterDigitalOceanCloudSpec
type ExternalClusterDigitalOceanCloudSpec struct {

	// Token is used to authenticate with the DigitalOcean API.
	Token string `json:"token,omitempty"`
}

// Validate validates this external cluster digital ocean cloud spec
func (m *ExternalClusterDigitalOceanCloudSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this external cluster digital ocean cloud spec based on context it is used
func (m *ExternalClusterDigitalOceanCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalClusterDigitalOceanCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalClusterDigitalOceanCloudSpec) UnmarshalBinary(b []byte) error {
	var res ExternalClusterDigitalOceanCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
