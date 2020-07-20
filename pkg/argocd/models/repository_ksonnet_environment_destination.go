// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RepositoryKsonnetEnvironmentDestination repository ksonnet environment destination
//
// swagger:model repositoryKsonnetEnvironmentDestination
type RepositoryKsonnetEnvironmentDestination struct {

	// Namespace is the namespace of the Kubernetes server that targets should be deployed to
	Namespace string `json:"namespace,omitempty"`

	// Server is the Kubernetes server that the cluster is running on.
	Server string `json:"server,omitempty"`
}

// Validate validates this repository ksonnet environment destination
func (m *RepositoryKsonnetEnvironmentDestination) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RepositoryKsonnetEnvironmentDestination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepositoryKsonnetEnvironmentDestination) UnmarshalBinary(b []byte) error {
	var res RepositoryKsonnetEnvironmentDestination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
