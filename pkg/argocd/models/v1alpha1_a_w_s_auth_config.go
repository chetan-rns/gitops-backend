// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1AWSAuthConfig AWSAuthConfig is an AWS IAM authentication configuration
//
// swagger:model v1alpha1AWSAuthConfig
type V1alpha1AWSAuthConfig struct {

	// ClusterName contains AWS cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// RoleARN contains optional role ARN. If set then AWS IAM Authenticator assume a role to perform cluster operations instead of the default AWS credential provider chain.
	RoleARN string `json:"roleARN,omitempty"`
}

// Validate validates this v1alpha1 a w s auth config
func (m *V1alpha1AWSAuthConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1AWSAuthConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1AWSAuthConfig) UnmarshalBinary(b []byte) error {
	var res V1alpha1AWSAuthConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
