// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ResourceRef ResourceRef includes fields which unique identify resource
//
// swagger:model v1alpha1ResourceRef
type V1alpha1ResourceRef struct {

	// group
	Group string `json:"group,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1alpha1 resource ref
func (m *V1alpha1ResourceRef) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ResourceRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ResourceRef) UnmarshalBinary(b []byte) error {
	var res V1alpha1ResourceRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
