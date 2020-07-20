// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1alpha1OrphanedResourcesMonitorSettings OrphanedResourcesMonitorSettings holds settings of orphaned resources monitoring
//
// swagger:model v1alpha1OrphanedResourcesMonitorSettings
type V1alpha1OrphanedResourcesMonitorSettings struct {

	// Warn indicates if warning condition should be created for apps which have orphaned resources
	Warn bool `json:"warn,omitempty"`
}

// Validate validates this v1alpha1 orphaned resources monitor settings
func (m *V1alpha1OrphanedResourcesMonitorSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1alpha1OrphanedResourcesMonitorSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1alpha1OrphanedResourcesMonitorSettings) UnmarshalBinary(b []byte) error {
	var res V1alpha1OrphanedResourcesMonitorSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
