// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImageSources ImageSources represents standard VM Image sources.
//
// swagger:model ImageSources
type ImageSources struct {

	// EnableCustomImages allows to enable/disable the usage of custom-disks (defaults to false).
	EnableCustomImages bool `json:"enableCustomImages,omitempty"`

	// http
	HTTP *HTTPSource `json:"http,omitempty"`
}

// Validate validates this image sources
func (m *ImageSources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageSources) validateHTTP(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {
		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this image sources based on the context it is used
func (m *ImageSources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageSources) contextValidateHTTP(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTP != nil {
		if err := m.HTTP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageSources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageSources) UnmarshalBinary(b []byte) error {
	var res ImageSources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
