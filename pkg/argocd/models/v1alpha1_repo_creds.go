// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1RepoCreds RepoCreds holds a repository credentials definition
//
// swagger:model v1alpha1RepoCreds
type V1alpha1RepoCreds struct {

	// Password for authenticating at the repo server
	Password string `json:"password,omitempty"`

	// SSH private key data for authenticating at the repo server (only Git repos)
	SSHPrivateKey string `json:"sshPrivateKey,omitempty"`

	// TLS client cert data for authenticating at the repo server
	TLSClientCertData string `json:"tlsClientCertData,omitempty"`

	// TLS client cert key for authenticating at the repo server
	TLSClientCertKey string `json:"tlsClientCertKey,omitempty"`

	// URL is the URL that this credentials matches to
	URL string `json:"url,omitempty"`

	// Username for authenticating at the repo server
	Username string `json:"username,omitempty"`
}

// Validate validates this v1alpha1 repo creds
func (m *V1alpha1RepoCreds) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1RepoCreds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1RepoCreds) UnmarshalBinary(b []byte) error {
	var res V1alpha1RepoCreds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
