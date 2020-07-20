// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1TLSClientConfig TLSClientConfig contains settings to enable transport layer security
//
// swagger:model v1alpha1TLSClientConfig
type V1alpha1TLSClientConfig struct {

	// CAData holds PEM-encoded bytes (typically read from a root certificates bundle).
	// CAData takes precedence over CAFile
	// Format: byte
	CaData strfmt.Base64 `json:"caData,omitempty"`

	// CertData holds PEM-encoded bytes (typically read from a client certificate file).
	// CertData takes precedence over CertFile
	// Format: byte
	CertData strfmt.Base64 `json:"certData,omitempty"`

	// Server should be accessed without verifying the TLS certificate. For testing only.
	Insecure bool `json:"insecure,omitempty"`

	// KeyData holds PEM-encoded bytes (typically read from a client certificate key file).
	// KeyData takes precedence over KeyFile
	// Format: byte
	KeyData strfmt.Base64 `json:"keyData,omitempty"`

	// ServerName is passed to the server for SNI and is used in the client to check server
	// certificates against. If ServerName is empty, the hostname used to contact the
	// server is used.
	ServerName string `json:"serverName,omitempty"`
}

// Validate validates this v1alpha1 TLS client config
func (m *V1alpha1TLSClientConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1TLSClientConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1TLSClientConfig) UnmarshalBinary(b []byte) error {
	var res V1alpha1TLSClientConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
