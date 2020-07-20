// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ApplicationSourceHelm ApplicationSourceHelm holds helm specific options
//
// swagger:model v1alpha1ApplicationSourceHelm
type V1alpha1ApplicationSourceHelm struct {

	// FileParameters are file parameters to the helm template
	FileParameters []*V1alpha1HelmFileParameter `json:"fileParameters"`

	// Parameters are parameters to the helm template
	Parameters []*V1alpha1HelmParameter `json:"parameters"`

	// The Helm release name. If omitted it will use the application name
	ReleaseName string `json:"releaseName,omitempty"`

	// ValuesFiles is a list of Helm value files to use when generating a template
	ValueFiles []string `json:"valueFiles"`

	// Values is Helm values, typically defined as a block
	Values string `json:"values,omitempty"`
}

// Validate validates this v1alpha1 application source helm
func (m *V1alpha1ApplicationSourceHelm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSourceHelm) validateFileParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.FileParameters) { // not required
		return nil
	}

	for i := 0; i < len(m.FileParameters); i++ {
		if swag.IsZero(m.FileParameters[i]) { // not required
			continue
		}

		if m.FileParameters[i] != nil {
			if err := m.FileParameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fileParameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1alpha1ApplicationSourceHelm) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ApplicationSourceHelm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationSourceHelm) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationSourceHelm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}