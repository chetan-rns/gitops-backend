// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1ApplicationSource ApplicationSource contains information about github repository, path within repository and target application environment.
//
// swagger:model v1alpha1ApplicationSource
type V1alpha1ApplicationSource struct {

	// Chart is a Helm chart name
	Chart string `json:"chart,omitempty"`

	// directory
	Directory *V1alpha1ApplicationSourceDirectory `json:"directory,omitempty"`

	// helm
	Helm *V1alpha1ApplicationSourceHelm `json:"helm,omitempty"`

	// ksonnet
	Ksonnet *V1alpha1ApplicationSourceKsonnet `json:"ksonnet,omitempty"`

	// kustomize
	Kustomize *V1alpha1ApplicationSourceKustomize `json:"kustomize,omitempty"`

	// Path is a directory path within the Git repository
	Path string `json:"path,omitempty"`

	// plugin
	Plugin *V1alpha1ApplicationSourcePlugin `json:"plugin,omitempty"`

	// RepoURL is the repository URL of the application manifests
	RepoURL string `json:"repoURL,omitempty"`

	// TargetRevision defines the commit, tag, or branch in which to sync the application to.
	// If omitted, will sync to HEAD
	TargetRevision string `json:"targetRevision,omitempty"`
}

// Validate validates this v1alpha1 application source
func (m *V1alpha1ApplicationSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHelm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKsonnet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKustomize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1alpha1ApplicationSource) validateDirectory(formats strfmt.Registry) error {

	if swag.IsZero(m.Directory) { // not required
		return nil
	}

	if m.Directory != nil {
		if err := m.Directory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("directory")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationSource) validateHelm(formats strfmt.Registry) error {

	if swag.IsZero(m.Helm) { // not required
		return nil
	}

	if m.Helm != nil {
		if err := m.Helm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationSource) validateKsonnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Ksonnet) { // not required
		return nil
	}

	if m.Ksonnet != nil {
		if err := m.Ksonnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ksonnet")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationSource) validateKustomize(formats strfmt.Registry) error {

	if swag.IsZero(m.Kustomize) { // not required
		return nil
	}

	if m.Kustomize != nil {
		if err := m.Kustomize.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kustomize")
			}
			return err
		}
	}

	return nil
}

func (m *V1alpha1ApplicationSource) validatePlugin(formats strfmt.Registry) error {

	if swag.IsZero(m.Plugin) { // not required
		return nil
	}

	if m.Plugin != nil {
		if err := m.Plugin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plugin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1ApplicationSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1ApplicationSource) UnmarshalBinary(b []byte) error {
	var res V1alpha1ApplicationSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
