// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HelmDeployOptions HelmDeployOptions holds the deployment settings when templating method is Helm.
//
// swagger:model HelmDeployOptions
type HelmDeployOptions struct {

	// Atomic corresponds to the --atomic flag on Helm cli.
	// if set, the installation process deletes the installation on failure; the upgrade process rolls back changes made in case of failed upgrade.
	Atomic bool `json:"atomic,omitempty"`

	// Wait corresponds to the --wait flag on Helm cli.
	// if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as timeout
	Wait bool `json:"wait,omitempty"`

	// timeout
	Timeout Duration `json:"timeout,omitempty"`
}

// Validate validates this helm deploy options
func (m *HelmDeployOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this helm deploy options based on context it is used
func (m *HelmDeployOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HelmDeployOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelmDeployOptions) UnmarshalBinary(b []byte) error {
	var res HelmDeployOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
