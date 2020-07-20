// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RepositoryKsonnetEnvironment repository ksonnet environment
//
// swagger:model repositoryKsonnetEnvironment
type RepositoryKsonnetEnvironment struct {

	// destination
	Destination *RepositoryKsonnetEnvironmentDestination `json:"destination,omitempty"`

	// KubernetesVersion is the kubernetes version the targeted cluster is running on.
	K8sVersion string `json:"k8sVersion,omitempty"`

	// Name is the user defined name of an environment
	Name string `json:"name,omitempty"`
}

// Validate validates this repository ksonnet environment
func (m *RepositoryKsonnetEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepositoryKsonnetEnvironment) validateDestination(formats strfmt.Registry) error {

	if swag.IsZero(m.Destination) { // not required
		return nil
	}

	if m.Destination != nil {
		if err := m.Destination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("destination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepositoryKsonnetEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepositoryKsonnetEnvironment) UnmarshalBinary(b []byte) error {
	var res RepositoryKsonnetEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
