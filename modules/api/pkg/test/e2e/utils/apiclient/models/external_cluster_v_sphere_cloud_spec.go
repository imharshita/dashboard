// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExternalClusterVSphereCloudSpec ExternalClusterVSphereCloudSpec credentials represents a credential for accessing vSphere.
//
// swagger:model ExternalClusterVSphereCloudSpec
type ExternalClusterVSphereCloudSpec struct {

	// password
	Password string `json:"password,omitempty"`

	// server
	Server string `json:"server,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this external cluster v sphere cloud spec
func (m *ExternalClusterVSphereCloudSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this external cluster v sphere cloud spec based on context it is used
func (m *ExternalClusterVSphereCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalClusterVSphereCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalClusterVSphereCloudSpec) UnmarshalBinary(b []byte) error {
	var res ExternalClusterVSphereCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
